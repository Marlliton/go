package exam

import "github.com/google/uuid"

type QuestionItem struct {
	id    string
	text  string
	right bool
}

func NewQuestionItem(id, text string, right bool) (*QuestionItem, error) {
	if text == "" {
		return nil, ErrMissingValues
	}
	if id == "" {
		id = uuid.New().String()
	} else {
		if _, err := uuid.Parse(id); err != nil {
			return nil, ErrInvalidId
		}
	}

	return &QuestionItem{
		id,
		text,
		right,
	}, nil
}
func (qi *QuestionItem) GetID() string    { return qi.id }
func (qi *QuestionItem) GetText() string  { return qi.text }
func (qi *QuestionItem) GetIsRight() bool { return qi.right }
