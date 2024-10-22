package mongorepo

import (
	"context"
	"testing"

	"github.com/Marlliton/go-quizzer/domain/exam"
)

func TestMongoRepository_Save(t *testing.T) {
	repo, err := New(context.Background(), "mongodb://localhost:27017")
	if err != nil {
		t.Fatal(err)
	}

	toSaveExam := init_exam(t)
	if err != nil {
		t.Fatal(err)
	}

	testeCases := []struct {
		name        string
		param       *exam.Exam
		exec        func(*exam.Exam) error
		expectedErr error
	}{
		{
			name:        "Save Exam",
			param:       toSaveExam,
			exec:        func(e *exam.Exam) error { return repo.Save(e) },
			expectedErr: nil,
		},
	}

	for _, tc := range testeCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.exec(tc.param)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
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
