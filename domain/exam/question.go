package exam

type Question struct {
	id        string
	statement string
	items     []*QuestionItem
}

func (q *Question) GetID() string             { return q.id }
func (q *Question) GetStatement() string      { return q.statement }
func (q *Question) GetItems() []*QuestionItem { return q.items }
