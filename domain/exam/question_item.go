package exam

import (
	"github.com/Marlliton/go-quizzer/domain/fail"
	"github.com/google/uuid"
)

type QuestionItem struct {
	id    string
	text  string
	right bool
}

var errItemCode = "QuestionItem"

func NewQuestionItem(id, text string, right bool) (*QuestionItem, error) {
	if text == "" {
		return nil, fail.WithValidationError(errItemCode, "title are required")
	}
	if id == "" {
		id = uuid.New().String()
	} else {
		if _, err := uuid.Parse(id); err != nil {
			return nil, fail.WithValidationError(errItemCode, "invalid id format")
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
