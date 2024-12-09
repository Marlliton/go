package examsvc

import (
	"github.com/Marlliton/go-quizzer/domain/exam"
	"github.com/Marlliton/go-quizzer/domain/fail"
)

func (es *ExamService) Save(e *exam.Exam) error {
	existingExam, err := es.repo.Get(e.GetID())
	if err == nil && existingExam != nil {
		return fail.WithAlreadyExistsError(errCodeExamSvc, "exam already exists")
	}

	return es.repo.Save(e)
}
