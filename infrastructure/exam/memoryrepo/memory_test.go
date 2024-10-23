package memoryrepo

import (
	"errors"
	"testing"

	"github.com/Marlliton/go-quizzer/domain/exam"
	"github.com/Marlliton/go-quizzer/domain/fail"
)

func TestMemoryExam_Save(t *testing.T) {

	repo := New()
	toSaveExam, err := exam.NewExam("", "Existing Exam", "Just testing", nil)
	if err != nil {
		t.Fatal(err)
	}

	type testCase struct {
		name  string
		param *exam.Exam
		exec  func(*exam.Exam, *testing.T)
	}

	testCases := []testCase{
		{
			name:  "Save exam",
			param: toSaveExam,
			exec: func(e *exam.Exam, t *testing.T) {
				err := repo.Save(e)
				if err != nil {
					t.Fatalf("expected err nil, got %v", err)
				}
			},
		}, {
			name:  "Failed to save existing exam",
			param: toSaveExam,
			exec: func(e *exam.Exam, t *testing.T) {
				var errAlreadyExists *fail.AlreadyExistsError

				err := repo.Save(e)
				if !errors.As(err, &errAlreadyExists) {
					t.Fatalf("expected error %v, got %v", errAlreadyExists, err)
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

func TestMemoryExam_Get(t *testing.T) {

	repo := New()
	existingExam, err := exam.NewExam("", "Existing Exam", "Just testing", nil)
	if err != nil {
		t.Fatal(err)
	}
	secondExam, err := exam.NewExam("", "Second Exam", "Just testing", nil)
	if err != nil {
		t.Fatal(err)
	}

	repo.Save(existingExam)
	repo.Save(secondExam)

	type testCase struct {
		name  string
		param string
		exec  func(string, *testing.T)
	}

	testCases := []testCase{
		{
			name:  "Get existing by id",
			param: existingExam.GetID(),
			exec: func(id string, t *testing.T) {
				t.Helper()

				_, err := repo.Get(id)

				if err != nil {
					t.Fatalf("expected err nil, got %v", err)
				}
			},
		}, {
			name:  "Get non-existing by id",
			param: "non-existing",
			exec: func(id string, t *testing.T) {
				t.Helper()
				var errNotFound *fail.NotFoundError

				_, err := repo.Get(id)
				if !errors.As(err, &errNotFound) {
					t.Fatalf("expected error %v, got %v", errNotFound, err)
				}
			},
		}, {
			name:  "Get all exams",
			param: "",
			exec: func(_ string, t *testing.T) {
				t.Helper()
				exams, err := repo.GetAll()
				if err != nil {
					t.Fatalf("expected error nil, got %v", err)
				}

				if len(exams) != 2 {
					t.Fatalf("expected at least tow exams, got %d", len(exams))
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

func TestMemoryExam_Update(t *testing.T) {
	repo := New()
	existingExam, err := exam.NewExam("", "Existing Exam", "Just testing", nil)
	if err != nil {
		t.Fatal(err)
	}
	nonExistingExam, err := exam.NewExam("", "Non-Existing Exam", "Just testing", nil)
	if err != nil {
		t.Fatal(err)
	}

	if err := repo.Save(existingExam); err != nil {
		t.Fatal(err)
	}

	type testCase struct {
		name  string
		param *exam.Exam
		exec  func(*exam.Exam, *testing.T)
	}

	testCases := []testCase{
		{
			name:  "Update exam title",
			param: existingExam,
			exec: func(e *exam.Exam, t *testing.T) {
				t.Helper()

				e.SetTitle("Title modifield")

				err := repo.Update(e)
				if err != nil {
					t.Fatalf("expected error nil, got %v", err)
				}
				updExam, err := repo.Get(e.GetID())
				if err != nil {
					return
				}
				if updExam.GetTitle() != "Title modifield" {
					t.Fatalf("expected title %s, got %s", "Title modifield", updExam.GetTitle())
				}
			},
		}, {
			name:  "Failed to save existing exam",
			param: nonExistingExam,
			exec: func(e *exam.Exam, t *testing.T) {
				t.Helper()
				var notFoundErr *fail.NotFoundError

				err := repo.Update(e)
				if !errors.As(err, &notFoundErr) {
					t.Fatalf("expected error %v, got %v", notFoundErr, err)

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

func TestMemoryExam_Delete(t *testing.T) {

	repo := New()
	exam1, err := exam.NewExam("", "Existing Exam", "Just testing", nil)
	if err != nil {
		t.Fatal(err)
	}
	exam2, err := exam.NewExam("", "Non-Existing Exam", "Just testing", nil)
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
		name  string
		param string
		exec  func(string, *testing.T)
	}

	testCases := []testCase{
		{
			name:  "Delete exam",
			param: exam1.GetID(),
			exec: func(id string, t *testing.T) {
				t.Helper()

				err := repo.Delete(id)
				if err != nil {
					t.Fatalf("expected err nil, got %v", err)
				}
			},
		}, {
			name:  "Failed to delete non-existing exam",
			param: "non-existing",
			exec: func(id string, t *testing.T) {
				t.Helper()

				var notFoundErr *fail.NotFoundError
				err := repo.Delete(id)

				if !errors.As(err, &notFoundErr) {
					t.Fatalf("expected error %v, got %v", notFoundErr, err)
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
