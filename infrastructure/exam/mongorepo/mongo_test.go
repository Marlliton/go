package mongorepo

import (
	"context"
	"errors"
	"log"
	"os"
	"testing"
	"time"

	"github.com/Marlliton/go-quizzer/domain/exam"
	"github.com/Marlliton/go-quizzer/domain/fail"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	ctx           = context.Background()
	uriConnection = "mongodb://localhost:27017"
	repo          *MongoRepository
)

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	repoCreated, err := New(ctx, uriConnection)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB %v", err)
	}

	repo = repoCreated

	code := m.Run()

	if err := repo.db.Drop(ctx); err != nil {
		log.Fatalf("failed to drop the database %v", err)
	}

	if err := repo.db.Client().Disconnect(ctx); err != nil {
		log.Fatalf("failed to disconnect from MongoDB %v", err)
	}

	os.Exit(code)
}

func TestMongoRepository_Get(t *testing.T) {

	testCases := []struct {
		name  string
		param string
		exec  func(string, *testing.T)
	}{
		{
			name: "Get exam by id",
			exec: func(_ string, t *testing.T) {
				t.Helper()

				savedExam := init_exam(t)

				err := repo.Save(savedExam)
				if err != nil {
					t.Fatalf("err saving exam %v", err)
				}

				e, err := repo.Get(savedExam.GetID())
				if err != nil {
					t.Fatalf("error getting exam %v", err)
				}

				if e.GetID() != savedExam.GetID() {
					t.Fatal("divergent result")
				}
			},
		}, {
			name:  "Get non-existing exam by id",
			param: "non-existing-id",
			exec: func(id string, t *testing.T) {
				t.Helper()
				var expectedErr *fail.NotFoundError

				_, err := repo.Get(id)
				if !errors.As(err, &expectedErr) {
					t.Fatalf("expectedErr: NotFoundError, got %v", err)
				}
			},
		}, {
			name: "Get all documents",
			exec: func(_ string, t *testing.T) {
				t.Helper()

				exam1 := init_exam(t)
				exam2 := init_exam(t)
				exam3 := init_exam(t)

				repo.Save(exam1)
				repo.Save(exam2)
				repo.Save(exam3)

				result, err := repo.GetAll()
				if err != nil {
					t.Fatalf("error getting all exams %v", err)
				}

				if len(result) != 3 {
					t.Fatalf("expected 3 documents, got %d", len(result))
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.exec(tc.param, t)
		})
		clearCollection(t)
	}
}

func TestMongoRepository_Save(t *testing.T) {
	toSaveExam := init_exam(t)

	testeCases := []struct {
		name  string
		param *exam.Exam
		exec  func(*exam.Exam, *testing.T)
	}{
		{
			name:  "Save Exam",
			param: toSaveExam,
			exec: func(e *exam.Exam, t *testing.T) {
				t.Helper()
				err := repo.Save(e)
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
			},
		},
	}

	for _, tc := range testeCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.exec(tc.param, t)
		})
		clearCollection(t)
	}
}

func TestMongoRepository_Update(t *testing.T) {
	testCases := []struct {
		name string
		exec func(t *testing.T)
	}{
		{
			name: "Update exam",
			exec: func(t *testing.T) {
				t.Helper()
				exam := init_exam(t)
				repo.Save(exam)

				exam.SetTitle("updated title")
				err := repo.Update(exam)
				if err != nil {
					t.Fatalf("error updating exam %v", err)
				}

				e, _ := repo.Get(exam.GetID())
				if e.GetTitle() != "updated title" {
					t.Fatalf(`expected title "updated title", got %v`, e.GetTitle())
				}
			},
		}, {
			name: "Update non-existing exam",
			exec: func(t *testing.T) {
				t.Helper()
				var expectedErr *fail.NotFoundError

				exam := init_exam(t)

				exam.SetTitle("updated title non-existing")
				_ = repo.Update(exam)

				_, err := repo.Get(exam.GetID())
				if !errors.As(err, &expectedErr) {
					t.Fatalf("expected error NotFoundError, got %v", err)
				}

			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.exec(t)
		})
		clearCollection(t)
	}
}

func TestMongoRepository_Delete(t *testing.T) {
	savedExam := init_exam(t)
	repo.Save(savedExam)

	testCases := []struct {
		name  string
		param string
		exec  func(string, *testing.T)
	}{
		{
			name:  "Delete exam by id",
			param: savedExam.GetID(),
			exec: func(id string, t *testing.T) {
				t.Helper()

				err := repo.Delete(id)
				if err != nil {
					t.Fatalf("failed to delete the exam: %s", id)
				}
			},
		}, {
			name:  "Delete non-existing exam by id",
			param: "non-existing-id",
			exec: func(id string, t *testing.T) {
				t.Helper()
				var expectedErr *fail.NotFoundError

				err := repo.Delete(id)
				if !errors.As(err, &expectedErr) {
					t.Fatalf("expected error NotFoundError, got %v", err)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.exec(tc.param, t)
		})
		clearCollection(t)
	}
}

func clearCollection(t *testing.T) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := repo.exam.DeleteMany(ctx, bson.M{})
	if err != nil {
		t.Fatalf("faile to clear collection exams %v", err)
	}
}

func init_exam(t *testing.T) *exam.Exam {
	t.Helper()

	item1, err := exam.NewQuestionItem("", "Opção A", true)
	if err != nil {
		t.Fatal(err)
	}

	item2, err := exam.NewQuestionItem("", "Opção B", false)
	if err != nil {
		t.Fatal(err)
	}

	question, err := exam.NewQuestion("", "Pergunta 01", []*exam.QuestionItem{item1, item2})
	if err != nil {
		t.Fatal(err)
	}

	ex, err := exam.NewExam("", "To Save", "Justo to save...", []*exam.Question{question})

	return ex
}
