package dto

type ExamDTOResponse struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Questions   []*QuestionDTOResponse `json:"questions"`
}

type QuestionDTOResponse struct {
	ID        string                     `json:"id"`
	Statement string                     `json:"statement"`
	Items     []*QuestionItemDTOResponse `json:"items"`
}

type QuestionItemDTOResponse struct {
	ID    string `json:"id"`
	Text  string `json:"text"`
	Right bool   `json:"right"`
}

type ExamDTORequest struct {
	Title       string               `json:"title" binding:"required"`
	Description string               `json:"description"`
	Questions   []QuestionDTORequest `json:"questions" binding:"dive"` // dive faz a validação recursiva no objetos filhos
}

type QuestionDTORequest struct {
	Statement string                   `json:"statement" binding:"required"`
	Items     []QuestionItemDTORequest `json:"items" binding:"dive"`
}

type QuestionItemDTORequest struct {
	Text  string `json:"text" binding:"required"`
	Right bool   `json:"right"`
}
