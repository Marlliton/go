package exam

import "github.com/google/uuid"

type Question struct {
	id        string
	statement string
	items     []*QuestionItem
}

func NewQuestion(id, statement string, items []*QuestionItem) (*Question, error) {
	if statement == "" {
		return nil, ErrMissingValues
	}
	if id == "" {
		id = uuid.New().String()
	} else {
		if _, err := uuid.Parse(id); err != nil {
			return nil, ErrInvalidId
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
