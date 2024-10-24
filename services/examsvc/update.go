package examsvc

import (
	"github.com/Marlliton/go-quizzer/domain/exam"
)

func (es *ExamService) Update(e *exam.Exam) error {
	_, err := es.repo.Get(e.GetID())
	if err != nil {
		return err
	}

	return es.repo.Update(e)
}
