package mapper

import (
	"github.com/Marlliton/go-quizzer/domain/exam"
	"github.com/Marlliton/go-quizzer/infra/api/dto"
)

func ToExamDomain(body *dto.ExamDTORequest) (*exam.Exam, error) {
	questions, err := toQuestionDomain(body.Questions)
	if err != nil {
		return nil, err
	}

	exam, err := exam.NewExam("", body.Title, body.Description, questions)
	if err != nil {
		return nil, err
	}

	return exam, nil
}

func toQuestionDomain(questions []dto.QuestionDTORequest) ([]*exam.Question, error) {
	size := len(questions)
	questionsDomain := make([]*exam.Question, size)

	for i, ques := range questions {
		items, err := toQuestionItemsDomain(ques.Items)
		if err != nil {
			return nil, err
		}
		question, err := exam.NewQuestion("", ques.Statement, items)
		if err != nil {
			return nil, err
		}

		questionsDomain[i] = question
	}

	return questionsDomain, nil
}

func toQuestionItemsDomain(items []dto.QuestionItemDTORequest) ([]*exam.QuestionItem, error) {
	size := len(items)
	itemsDomain := make([]*exam.QuestionItem, size)

	for i, it := range items {
		item, err := exam.NewQuestionItem("", it.Text, it.Right)
		if err != nil {
			return nil, err
		}

		itemsDomain[i] = item
	}

	return itemsDomain, nil
}
func ToExamDTO(e exam.Exam) *dto.ExamDTOResponse {
	return &dto.ExamDTOResponse{
		ID:          e.GetID(),
		Title:       e.GetTitle(),
		Description: e.GetDescription(),
		Questions:   toQuestionDTOResponse(e.GetQuestions()),
	}
}

func toQuestionDTOResponse(questions []*exam.Question) []*dto.QuestionDTOResponse {
	result := make([]*dto.QuestionDTOResponse, len(questions))

	for i, q := range questions {
		items := toQuestionItemDTOResponse(q.GetItems())
		result[i] = &dto.QuestionDTOResponse{
			ID:        q.GetID(),
			Statement: q.GetStatement(),
			Items:     items,
		}
	}

	return result
}

func toQuestionItemDTOResponse(items []*exam.QuestionItem) []*dto.QuestionItemDTOResponse {
	result := make([]*dto.QuestionItemDTOResponse, len(items))

	for i, it := range items {
		result[i] = &dto.QuestionItemDTOResponse{
			ID:    it.GetID(),
			Text:  it.GetText(),
			Right: it.GetIsRight(),
		}
	}

	return result
}
