package memoryrepo

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
func (mr *MemoryRepository) Save(addedExam *exam.Exam) error {
	mr.Lock()
	defer mr.Unlock()
	if _, ok := mr.exams[addedExam.GetID()]; ok {
		return exam.ErrExamAlreadyExists
	}

	mr.exams[addedExam.GetID()] = addedExam
	return nil
}
func (mr *MemoryRepository) Update(updExam *exam.Exam) error {
	mr.Lock()
	defer mr.Unlock()
	if _, ok := mr.exams[updExam.GetID()]; !ok {
		return exam.ErrExamNotFound
	}
	mr.exams[updExam.GetID()] = updExam
	return nil
}
func (mr *MemoryRepository) Delete(id string) error {
	mr.Lock()
	defer mr.Unlock()
	if _, ok := mr.exams[id]; !ok {
		return exam.ErrExamNotFound
	}

	delete(mr.exams, id)
	return nil
}
