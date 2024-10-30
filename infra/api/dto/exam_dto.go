package dto

import "github.com/Marlliton/go-quizzer/domain/exam"

type ExamDTO struct {
	Entity *exam.Exam
}

type examDTOResponse struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Questions   []*questionDTOResponse `json:"questions"`
}

type questionDTOResponse struct {
	ID        string                     `json:"id"`
	Statement string                     `json:"statement"`
	Items     []*questionItemDTOResponse `json:"items"`
}

type questionItemDTOResponse struct {
	ID    string `json:"id"`
	Text  string `json:"text"`
	Right bool   `json:"right"`
}

func (ed *ExamDTO) ToResponse() *examDTOResponse {
	return &examDTOResponse{
		ID:          ed.Entity.GetID(),
		Title:       ed.Entity.GetTitle(),
		Description: ed.Entity.GetDescription(),
		Questions:   toQuestionDTOResponse(ed.Entity.GetQuestions()),
	}
}

func toQuestionDTOResponse(questions []*exam.Question) []*questionDTOResponse {
	result := make([]*questionDTOResponse, len(questions))

	for i, q := range questions {
		items := toQuestionItemDTOResponse(q.GetItems())
		result[i] = &questionDTOResponse{
			ID:        q.GetID(),
			Statement: q.GetStatement(),
			Items:     items,
		}
	}

	return result
}

func toQuestionItemDTOResponse(items []*exam.QuestionItem) []*questionItemDTOResponse {
	result := make([]*questionItemDTOResponse, len(items))

	for i, it := range items {
		result[i] = &questionItemDTOResponse{
			ID:    it.GetID(),
			Text:  it.GetText(),
			Right: it.GetIsRight(),
		}
	}

	return result
}
