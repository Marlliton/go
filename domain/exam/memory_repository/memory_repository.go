package memoryrepository

import (
	"sync"

	"github.com/Marlliton/go-quizzer/domain/exam"
)

type MemoryRepository struct {
	exams map[string]*exam.Exam
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		exams: make(map[string]*exam.Exam),
	}
}

func (mr *MemoryRepository) Get(id string) (*exam.Exam, error) {
	if e, ok := mr.exams[id]; ok {
		return e, nil
	}

	return nil, exam.ErrExamNotFound
}
func (mr *MemoryRepository) GetAll() ([]*exam.Exam, error) {
	var exams []*exam.Exam

	for _, ex := range mr.exams {
		exams = append(exams, ex)
	}
	return exams, nil
}
func (mr *MemoryRepository) Save(*exam.Exam) error {
	return nil
}
func (mr *MemoryRepository) Update(*exam.Exam) error {
	return nil
}
func (mr *MemoryRepository) Delete(string) error {
	return nil
}
