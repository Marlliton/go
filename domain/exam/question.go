package exam

import (
	"github.com/Marlliton/go-quizzer/domain/fail"
	"github.com/google/uuid"
)

type Question struct {
	id        string
	statement string
	items     []*QuestionItem
}

var errQuestionCode = "Question"

func NewQuestion(id, statement string, items []*QuestionItem) (*Question, error) {
	if statement == "" {
		return nil, fail.WithValidationError(errQuestionCode, "statement are required")
	}
	if id == "" {
		id = uuid.New().String()
	} else {
		if _, err := uuid.Parse(id); err != nil {
			return nil, fail.WithValidationError(errQuestionCode, "invalid id format")
		}
	}

	return &Question{
		id,
		statement,
		items,
	}, nil
}

func (q *Question) GetID() string             { return q.id }
func (q *Question) GetStatement() string      { return q.statement }
func (q *Question) GetItems() []*QuestionItem { return q.items }
func (q *Question) GetCorrectItem(id string) *QuestionItem {
	for _, item := range q.items {
		if item.id == id {
			return item
		}
	}

	return nil
}
