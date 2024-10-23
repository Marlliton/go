package examsvc

import (
	"github.com/Marlliton/go-quizzer/domain/exam"
	"github.com/Marlliton/go-quizzer/domain/fail"
	"github.com/Marlliton/go-quizzer/infrastructure/exam/memoryrepo"
)

var errCodeExamSvc = "ExamService"

type ExamServiceConfig func(es *ExamService) error

type ExamService struct {
	repo exam.Repository
}

func NewExamsvc(cfgs ...ExamServiceConfig) (*ExamService, error) {
	es := &ExamService{}

	for _, cfg := range cfgs {
		err := cfg(es)
		if err != nil {
			return nil, err
		}
	}

	return es, nil
}

func (es *ExamService) Save(e *exam.Exam) error {
	existingExam, err := es.repo.Get(e.GetID())
	if err == nil && existingExam != nil {
		return fail.WithAlreadyExistsError(errCodeExamSvc, "exam already exists")
	}

	err = es.repo.Save(e)
	if err != nil {
		return err
	}

	return nil
}

func WithMemoryExamRepository() ExamServiceConfig {
	repo := memoryrepo.New()

	return func(es *ExamService) error {
		es.repo = repo
		return nil
	}
}
