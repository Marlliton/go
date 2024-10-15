package exam

type Exam struct {
	id          string
	title       string
	description string
	questions   []*Question
}

func (e *Exam) GetID() string            { return e.id }
func (e *Exam) GetTitle() string         { return e.title }
func (e *Exam) GetDescription() string   { return e.description }
func (e *Exam) GetQuestion() []*Question { return e.questions }
