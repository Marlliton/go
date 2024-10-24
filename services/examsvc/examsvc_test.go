package examsvc

import (
	"context"
	"testing"

	"github.com/Marlliton/go-quizzer/domain/exam"
)

func TestExamService_AllMethods(t *testing.T) {

	svc, err := NewExamsvc(
		// WithMemoryExamRepository(),
		WithMongoExamRepository(context.Background(), "mongodb://localhost:27017"),
	)
	if err != nil {
		t.Fatalf("error creating exam service %v", err)
	}

	testCases := []struct {
		name string
		exec func(*testing.T)
	}{
		{
			name: "Save exam",
			exec: func(t *testing.T) {
				t.Helper()
				exam := init_exam(t)

				err := svc.Save(exam)
				if err != nil {
					t.Fatalf("error saving exam %v", err)
				}

				_ = svc.Delete(exam.GetID())
			},
		}, {
			name: "Get exam by id",
			exec: func(t *testing.T) {
				t.Helper()
				exam := init_exam(t)

				_ = svc.Save(exam)

				_, err := svc.Get(exam.GetID())
				if err != nil {
					t.Fatalf("error getting exam by id %v", err)
				}

				_ = svc.Delete(exam.GetID())
			},
		}, {
			name: "Get all exams",
			exec: func(t *testing.T) {
				t.Helper()
				exam1 := init_exam(t)
				exam2 := init_exam(t)

				_ = svc.Save(exam1)
				_ = svc.Save(exam2)

				result, err := svc.GetAll()
				if err != nil {
					t.Fatalf("error getting all exams %v", err)
				}

				if len(result) != 2 {
					t.Fatalf("expecte 2 results, got %d", len(result))
				}
				_ = svc.Delete(exam1.GetID())
				_ = svc.Delete(exam2.GetID())
			},
		}, {
			name: "Update exam",
			exec: func(t *testing.T) {
				t.Helper()
				exam := init_exam(t)

				_ = svc.Save(exam)
				newTitle := "updated title"
				exam.SetTitle(newTitle)

				err := svc.Update(exam)
				if err != nil {
					t.Fatalf("error updating exam %v", err)
				}
				_ = svc.Delete(exam.GetID())
			},
		}, {
			name: "Delete exam",
			exec: func(t *testing.T) {
				t.Helper()
				exam := init_exam(t)

				_ = svc.Save(exam)

				err := svc.Delete(exam.GetID())
				if err != nil {
					t.Fatalf("error deleting exam %v", err)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.exec(t)
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
