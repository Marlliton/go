package mongorepo

import (
	"context"
	"errors"
	"testing"

	"github.com/Marlliton/go-quizzer/domain/exam"
	"github.com/Marlliton/go-quizzer/domain/fail"
)

var (
	ctx           = context.Background()
	uriConnection = "mongodb://localhost:27017"
)

func TestMongoRepository_Get(t *testing.T) {
	repo, err := New(ctx, uriConnection)
	if err != nil {
		t.Fatal(err)
	}

	savedExam := init_exam(t)
	err = repo.Save(savedExam)
	if err != nil {
		t.Fatalf("err saving exam %v", err)
	}

	testCases := []struct {
		name  string
		param string
		exec  func(string, *testing.T)
	}{
		{
			name:  "Get exam by id",
			param: savedExam.GetID(),
			exec: func(id string, t *testing.T) {
				t.Helper()
				e, err := repo.Get(id)
				if err != nil {
					t.Fatalf("error getting exam %v", err)
				}

				if e.GetID() != id {
					t.Fatal("divergent result")
				}
			},
		}, {
			name:  "No exam result by id",
			param: "Get non-existing exam by id",
			exec: func(id string, t *testing.T) {
				t.Helper()
				var expectedErr *fail.NotFoundError

				_, err := repo.Get(id)
				if !errors.As(err, &expectedErr) {
					t.Fatalf("expectedErr: NotFoundError, got %v", err)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.exec(tc.param, t)
		})
	}
}

func TestMongoRepository_Save(t *testing.T) {
	repo, err := New(ctx, uriConnection)
	if err != nil {

		t.Fatal(err)
	}

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
