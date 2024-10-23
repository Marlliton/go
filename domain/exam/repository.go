package exam

type Repository interface {
	Get(string) (*Exam, error)
	GetAll() ([]*Exam, error)
	Save(*Exam) error
	Update(*Exam) error
	Delete(string) error
}
