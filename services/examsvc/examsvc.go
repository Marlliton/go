package examsvc

import (
	"context"

	"github.com/Marlliton/go-quizzer/domain/exam"
	"github.com/Marlliton/go-quizzer/infra/database/memory"
	"github.com/Marlliton/go-quizzer/infra/database/mongo"
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

func WithMongoExamRepository(ctx context.Context, uriConnection string) ExamServiceConfig {
	repo, err := mongo.NewMongoExamRepository(ctx, uriConnection)

	return func(es *ExamService) error {
		if err != nil {
			return err
		}

		es.repo = repo

		return nil
	}
}

func WithMemoryExamRepository() ExamServiceConfig {
	repo := memory.NewInMemoryExamRepository()

	return func(es *ExamService) error {
		es.repo = repo
		return nil
	}
}
