package exam

type answer struct {
	questionID     string
	questionItemID string
	right          bool
}

type ExamAnswer struct {
	examID      string
	responderID string
	exam        *Exam
	answers     []*answer
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

func (ea *ExamAnswer) ReplyWith(questionID, itemID string) {
	question := ea.getQuestionByID(questionID)
	item := question.GetCorrectItem(itemID)

	// TODO: Verify if question has already been answered
	answers := append(ea.answers, &answer{
		questionID:     questionID,
		questionItemID: itemID,
		right:          item.GetIsRight(),
	})
	ea.answers = answers
}

// func (ea *ExamAnswer) GetAccuracyRate() float32 {
// 	questions := ea.getQuestions()
//
// }
