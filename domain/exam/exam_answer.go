package exam

import (
	"errors"
)

var (
	ErrQuestionNotFound = errors.New("question not found")
	ErrAnswerNotFound   = errors.New("answer not found")
)

type Answer struct {
	questionID     string
	questionItemID string
	right          bool
}

type ExamAnswer struct {
	examID            string
	responderID       string
	currentQuestionID string
	exam              *Exam
	answers           []*Answer
}

func (ea *ExamAnswer) getQuestions() []*Question {
	return ea.exam.GetQuestions()
}
func (ea *ExamAnswer) getQuestionByID(id string) *Question {
	for _, ques := range ea.getQuestions() {
		if ques.GetID() == id {
			return ques
		}
	}

	return nil
}

func (ea *ExamAnswer) SubmitAnswer(questionID, itemID string) error {
	question := ea.getQuestionByID(questionID)
	if question == nil {
		return ErrQuestionNotFound
	}
	item := question.GetCorrectItem(itemID)
	if item == nil {
		return ErrAnswerNotFound
	}

	ea.currentQuestionID = questionID

	for _, answer := range ea.answers {
		if answer.questionID == questionID {

			answer.questionItemID = itemID
			answer.right = item.GetIsRight()
			return nil
		}
	}

	answers := append(ea.answers, &Answer{
		questionID:     questionID,
		questionItemID: itemID,
		right:          item.GetIsRight(),
	})
	ea.answers = answers

	return nil
}

func (ea *ExamAnswer) NextQuestion() {
	if ea.currentQuestionID == "" {
		return
	}

	questions := ea.getQuestions()
	if len(questions) == 0 {
		return
	}
	for i, ques := range questions {
		if ques.id == ea.currentQuestionID && i+1 < len(questions) {
			ea.currentQuestionID = questions[i+1].GetID()
			return
		}
	}
}

func (ea *ExamAnswer) PreviousQuestion() {
	if ea.currentQuestionID == "" {
		return
	}

	questions := ea.getQuestions()
	if len(questions) == 0 {
		return
	}

	for i, ques := range questions {
		if ques.id == ea.currentQuestionID && i-1 >= 0 {
			ea.currentQuestionID = questions[i-1].GetID()
			return
		}
	}

}

func (ea *ExamAnswer) GetCurrentQuestion() (*Question, error) {
	if ea.currentQuestionID == "" {
		if len(ea.getQuestions()) <= 0 {
			return nil, ErrQuestionNotFound
		}
		return ea.getQuestions()[0], nil
	}

	for _, ques := range ea.getQuestions() {
		if ques.id == ea.currentQuestionID {
			return ques, nil
		}
	}

	return nil, ErrQuestionNotFound
}

func (ea *ExamAnswer) Score() int {
	totalQuestion := ea.exam.GetTotalQuestions()
	var correctAnswers []*Answer

	for _, ans := range ea.answers {
		if ans.right {
			correctAnswers = append(correctAnswers, ans)
		}
	}

	totalScore := (float64(len(correctAnswers)) / float64(totalQuestion)) * 100
	return int(totalScore)
}
