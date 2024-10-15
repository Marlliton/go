package memory_test

import (
	"testing"

	"github.com/Marlliton/go-quizzer/domain/exam"
	"github.com/Marlliton/go-quizzer/domain/exam/memory"
)

func TestMemoryExam_Save(t *testing.T) {
	repo := memory.New()
	toSaveExam, err := exam.New("", "Existing Exam", "Just testing", nil)
	if err != nil {
		t.Fatal(err)
	}

	type testCase struct {
		name        string
		param       *exam.Exam
		exec        func(*exam.Exam) error
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Save exam",
			param:       toSaveExam,
			exec:        func(e *exam.Exam) error { return repo.Save(e) },
			expectedErr: nil,
		}, {
			name:        "Failed to save existing exam",
			param:       toSaveExam,
			exec:        func(e *exam.Exam) error { return repo.Save(e) },
			expectedErr: exam.ErrExamAlreadyExists,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.exec(tc.param)
			if err != tc.expectedErr {
				t.Fatalf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemoryExam_Get(t *testing.T) {
	repo := memory.New()
	existingExam, err := exam.New("", "Existing Exam", "Just testing", nil)
	if err != nil {
		t.Fatal(err)
	}
	secondExam, err := exam.New("", "Second Exam", "Just testing", nil)
	if err != nil {
		t.Fatal(err)
	}

	repo.Save(existingExam)
	repo.Save(secondExam)

	type testCase struct {
		name        string
		param       string
		exec        func(string) (interface{}, error)
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Get existing by id",
			param:       existingExam.GetID(),
			exec:        func(id string) (interface{}, error) { return repo.Get(id) },
			expectedErr: nil,
		}, {
			name:        "Get non-existing by id",
			param:       "non-existing",
			exec:        func(id string) (interface{}, error) { return repo.Get(id) },
			expectedErr: exam.ErrExamNotFound,
		}, {
			name:        "Get all exams",
			param:       "",
			exec:        func(_ string) (interface{}, error) { return repo.GetAll() },
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.exec(tc.param)
			if err != tc.expectedErr {
				t.Fatalf("expected error %v, got %v", tc.expectedErr, err)
			}

			if tc.name == "Get all exams" {
				exams := result.([]*exam.Exam)

				if len(exams) != 2 {
					t.Fatalf("expected at least tow exams, got %d", len(exams))
				}
			}
		})
	}
}

func TestMemoryExam_Update(t *testing.T) {
	repo := memory.New()
	existingExam, err := exam.New("", "Existing Exam", "Just testing", nil)
	if err != nil {
		t.Fatal(err)
	}
	nonExistingExam, err := exam.New("", "Non-Existing Exam", "Just testing", nil)
	if err != nil {
		t.Fatal(err)
	}

	if err := repo.Save(existingExam); err != nil {
		t.Fatal(err)
	}

	type testCase struct {
		name        string
		param       *exam.Exam
		exec        func(*exam.Exam) error
		expectedErr error
	}

	testCases := []testCase{
		{
			name:  "Update exam title",
			param: existingExam,
			exec: func(e *exam.Exam) error {
				e.SetTitle("Title modifield")
				return repo.Update(e)
			},
			expectedErr: nil,
		}, {
			name:        "Failed to save existing exam",
			param:       nonExistingExam,
			exec:        func(e *exam.Exam) error { return repo.Update(e) },
			expectedErr: exam.ErrExamNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.exec(tc.param)
			if err != tc.expectedErr {
				t.Fatalf("expected error %v, got %v", tc.expectedErr, err)
			}

			updExam, err := repo.Get(tc.param.GetID())
			if err != nil {
				return
			}
			if updExam.GetTitle() != "Title modifield" {
				t.Fatalf("expected title %s, got %s", "Title modifield", updExam.GetTitle())
			}
		})
	}
}

func TestMemoryExam_Delete(t *testing.T) {
	repo := memory.New()
	exam1, err := exam.New("", "Existing Exam", "Just testing", nil)
	if err != nil {
		t.Fatal(err)
	}
	exam2, err := exam.New("", "Non-Existing Exam", "Just testing", nil)
	if err != nil {
		t.Fatal(err)
	}

	if err := repo.Save(exam1); err != nil {
		t.Fatal(err)
	}
	if err := repo.Save(exam2); err != nil {
		t.Fatal(err)
	}

	type testCase struct {
		name        string
		param       string
		exec        func(string) error
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Delete exam",
			param:       exam1.GetID(),
			exec:        func(id string) error { return repo.Delete(id) },
			expectedErr: nil,
		}, {
			name:        "Failed to delete non-existing exam",
			param:       "non-existing",
			exec:        func(id string) error { return repo.Delete(id) },
			expectedErr: exam.ErrExamNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.exec(tc.param)
			if err != tc.expectedErr {
				t.Fatalf("expected error %v, got %v", tc.expectedErr, err)
			}

			if tc.name == "Get all exams" {

				result, err := repo.GetAll()
				if err != nil {
					t.Fatal(err)
				}
				if len(result) != 1 {
					t.Fatalf("expected at least one exam, got %d", len(result))
				}
			}
		})
	}
}
