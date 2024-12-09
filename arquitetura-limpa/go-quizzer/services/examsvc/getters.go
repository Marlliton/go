package examsvc

import "github.com/Marlliton/go-quizzer/domain/exam"

func (es *ExamService) Get(id string) (*exam.Exam, error) {
	return es.repo.Get(id)
}

func (es *ExamService) GetAll() ([]*exam.Exam, error) {
	return es.repo.GetAll()
}
