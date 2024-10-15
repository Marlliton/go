package exam

type QuestionItem struct {
	id    string
	text  string
	right bool
}

func (qi *QuestionItem) GetID() string   { return qi.id }
func (qi *QuestionItem) GetText() string { return qi.text }
func (qi *QuestionItem) GetRight() bool  { return qi.right }
