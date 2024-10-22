package exam

import (
	"errors"

	"github.com/google/uuid"
)

// TODO: Padronizar todos os erros da aplicação
var (
	ErrInvalidId     = errors.New("invalid id format")
	ErrMissingValues = errors.New("missing values: title and description are required")
)

type Exam struct {
	id          string
	title       string
	description string
	questions   []*Question
}

func NewExam(id, title, description string, questions []*Question) (*Exam, error) {
	if title == "" || description == "" {
		return nil, ErrMissingValues
	}

	if id == "" {
		id = uuid.New().String()
	} else {
		if _, err := uuid.Parse(id); err != nil {
			return nil, ErrInvalidId
		}
	}

	newExam := &Exam{
		id:          id,
		title:       title,
		description: description,
		questions:   questions,
	}

	return newExam, nil
}

func (e *Exam) GetID() string             { return e.id }
func (e *Exam) GetTitle() string          { return e.title }
func (e *Exam) GetDescription() string    { return e.description }
func (e *Exam) GetQuestions() []*Question { return e.questions }
func (e *Exam) GetTotalQuestions() int    { return len(e.questions) }

func (e *Exam) SetTitle(t string)          { e.title = t }
func (e *Exam) SetDescription(d string)    { e.description = d }
func (e *Exam) SetQuestions(q []*Question) { e.questions = q }
