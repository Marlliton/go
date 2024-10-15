package exam

import "errors"

var (
	ErrExamNotFound      = errors.New("the exam was not found")
	ErrExamAlreadyExists = errors.New("the exam already exists")
)

type Repository interface {
	Get(string) (*Exam, error)
	GetAll() ([]*Exam, error)
	Save(*Exam) error
	Update(*Exam) error
	Delete(string) error
}
