package exam

import (
	"errors"
	"testing"

	"github.com/Marlliton/go-quizzer/domain/fail"
)

func TestExamAnswer_NextQuestion_And_PreviousQuestion(t *testing.T) {
	examAnswer := initExamAnswer()
	testeCases := []struct {
		name           string
		questionID     string
		itemID         string
		expectedResult string
	}{
		{
			name:           "Resposta incorreta para a pergunta 1",
			questionID:     "ques1",
			itemID:         "item1", // Resposta incorreta: "Um framework"
			expectedResult: "ques1",
		},
		{
			name:           "Resposta correta para a pergunta 2",
			questionID:     "ques2",
			itemID:         "item3", // A resposta correta é "Google"
			expectedResult: "ques2",
		},
	}
	err := examAnswer.SubmitAnswer(testeCases[0].questionID, testeCases[0].itemID)
	if err != nil {
		t.Fatal("Erro on SubmitAnswer", err)
	}

	t.Run(testeCases[0].name, func(t *testing.T) {

		if examAnswer.currentQuestionID != testeCases[0].expectedResult {
			t.Fatalf("expected result %s, got %s", testeCases[0].expectedResult, examAnswer.currentQuestionID)
		}

		examAnswer.NextQuestion()

		next, err := examAnswer.GetCurrentQuestion()
		if err != nil {
			t.Fatal(err)
		}

		if examAnswer.currentQuestionID != testeCases[1].expectedResult {
			t.Fatalf("expected result %s, got %s", testeCases[1].expectedResult, next.GetID())
		}
	})

	t.Run(testeCases[1].name, func(t *testing.T) {
		examAnswer.NextQuestion()

		next, err := examAnswer.GetCurrentQuestion()
		if err != nil {
			t.Fatal(err)
		}

		if examAnswer.currentQuestionID != testeCases[1].expectedResult {
			t.Fatalf("expected result %s, got %s", testeCases[1].expectedResult, next.GetID())
		}
	})

	// Test Previous Question
	t.Run(testeCases[1].name, func(t *testing.T) {
		examAnswer.PreviousQuestion()

		prev, err := examAnswer.GetCurrentQuestion()
		if err != nil {
			t.Fatal(err)
		}

		if examAnswer.currentQuestionID != testeCases[0].expectedResult {
			t.Fatalf("expected result %s, got %s", testeCases[0].expectedResult, prev.GetID())
		}
	})

	t.Run(testeCases[1].name, func(t *testing.T) {
		examAnswer.PreviousQuestion()

		prev, err := examAnswer.GetCurrentQuestion()
		if err != nil {
			t.Fatal(err)
		}

		if examAnswer.currentQuestionID != testeCases[0].expectedResult {
			t.Fatalf("expected result %s, got %s", testeCases[0].expectedResult, prev.GetID())
		}
	})
}

func TestExamAnswer_Score(t *testing.T) {
	examAnswer := initExamAnswer()
	testeCases := []struct {
		name        string
		questionID  string
		itemID      string
		expectedErr error
	}{
		{
			name:        "Resposta incorreta para a pergunta 1",
			questionID:  "ques1",
			itemID:      "item1", // Resposta incorreta: "Um framework"
			expectedErr: nil,
		},
		{
			name:        "Resposta correta para a pergunta 2",
			questionID:  "ques2",
			itemID:      "item3", // A resposta correta é "Google"
			expectedErr: nil,
		},
	}

	for _, tc := range testeCases {

		t.Run(tc.name, func(t *testing.T) {
			err := examAnswer.SubmitAnswer(tc.questionID, tc.itemID)
			if err != tc.expectedErr {
				t.Fatalf("Expected erro %v, got %v", tc.expectedErr, err)
			}
		})
	}

	score := examAnswer.Score()
	if score != 50 {
		t.Fatalf("expected socre %d, got %d", 50, score)
	}
}

func TestExamAnswer_SubmitQuestion(t *testing.T) {
	examAnswer := initExamAnswer()
	testeCases := []struct {
		name        string
		questionID  string
		itemID      string
		expectSave  bool // Define se esperamos que a resposta seja salva ou modificada
		isRight     bool
		expectedErr error
	}{
		{
			name:        "Resposta correta para a pergunta 1",
			questionID:  "ques1",
			itemID:      "item2", // A resposta correta é "Uma linguagem de programação"
			expectSave:  true,
			isRight:     true,
			expectedErr: nil,
		},
		{
			name:        "Resposta incorreta para a pergunta 1",
			questionID:  "ques1",
			itemID:      "item1", // Resposta incorreta: "Um framework"
			expectSave:  true,
			isRight:     false,
			expectedErr: nil,
		},
		{
			name:        "Resposta correta para a pergunta 2",
			questionID:  "ques2",
			itemID:      "item3", // A resposta correta é "Google"
			expectSave:  true,
			isRight:     true,
			expectedErr: nil,
		},
		{
			name:        "Submissão de pergunta inexistente",
			questionID:  "ques3", // Pergunta não existente
			itemID:      "item1",
			expectSave:  false, // Não deve adicionar nada
			isRight:     false,
			expectedErr: nil,
		},
		{
			name:        "Submissão duplicada para a mesma pergunta",
			questionID:  "ques1",
			itemID:      "item2", // Resposta já enviada
			expectSave:  true,    // Deve substituir a resposta anterior
			isRight:     true,
			expectedErr: nil,
		},
	}

	for _, tc := range testeCases {

		t.Run(tc.name, func(t *testing.T) {
			var errNotFound *fail.NotFoundError
			err := examAnswer.SubmitAnswer(tc.questionID, tc.itemID)

			if tc.name == "Submissão de pergunta inexistente" {
				if !errors.As(err, &errNotFound) {
					t.Fatalf("expected error %v, got %v", errNotFound, err)
				}
				return
			}

			if err != tc.expectedErr {
				t.Fatalf("Expected erro %v, got %v", tc.expectedErr, err)
			}

			if tc.expectSave {
				found := false

				for _, ans := range examAnswer.answers {
					if tc.questionID == ans.questionID && tc.itemID == ans.questionItemID {
						found = true
						if ans.right != tc.isRight {
							t.Fatalf("expected result for answer %t, got %t", tc.isRight, ans.right)
						}
						break
					}
				}

				if !found {
					t.Fatalf("answer not stored correctly for questionID %s", tc.questionID)
				}
			}
		})
	}
}

func initExamAnswer() *ExamAnswer {
	return &ExamAnswer{
		exam: &Exam{
			id:          "exam-id",
			title:       "Teste response",
			description: "Just a teste response",
			questions: []*Question{
				{
					id:        "ques1",
					statement: "O que é a Golang?",
					items: []*QuestionItem{
						{id: "item1", text: "Um framework", right: false},
						{id: "item2", text: "Uma linguagem de programação", right: true},
						{id: "item3", text: "Um sistema operacional", right: false},
						{id: "item4", text: "Uma biblioteca", right: false},
					},
				},
				{
					id:        "ques2",
					statement: "Quem desenvolveu o Golang?",
					items: []*QuestionItem{
						{id: "item1", text: "Microsoft", right: false},
						{id: "item2", text: "Apple", right: false},
						{id: "item3", text: "Google", right: true},
						{id: "item4", text: "Facebook", right: false},
					},
				},
			},
		},
		answers: []*Answer{},
	}
}
