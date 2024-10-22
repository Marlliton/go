package mongorepo

import (
	"context"
	"fmt"
	"time"

	"github.com/Marlliton/go-quizzer/domain/exam"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db   *mongo.Database
	exam *mongo.Collection
}

type mongoExam struct {
	id          string           `bson:"id"`
	title       string           `bson:"title"`
	description string           `bson:"description"`
	questions   []*mongoQuestion `bson:"questions"`
}

type mongoQuestion struct {
	id        string               `bson:"id"`
	statement string               `bson:"statement"`
	items     []*mongoQuestionItem `bson:"items"`
}

type mongoQuestionItem struct {
	id    string `bson:"id"`
	text  string `bson:"text"`
	right bool   `bson:"right"`
}

func (me *mongoExam) toAggregate() (*exam.Exam, error) {
	var questions []*exam.Question
	for _, mq := range me.questions {
		questionItems, err := me.buildQuestionItems(mq)
		if err != nil {
			return nil, fmt.Errorf("erro creating question items %s, %w", mq.id, err)
		}
		ques, err := exam.NewQuestion(mq.id, mq.statement, questionItems)
		if err != nil {
			return nil, fmt.Errorf("erro creating question %v", err)
		}

		questions = append(questions, ques)
	}

	return exam.NewExam(me.id, me.title, me.description, questions)
}
func (*mongoExam) buildQuestionItems(mq *mongoQuestion) ([]*exam.QuestionItem, error) {
	questionItems := make([]*exam.QuestionItem, len(mq.items))
	for i, mqi := range mq.items {
		item, err := exam.NewQuestionItem(mqi.id, mqi.text, mqi.right)
		if err != nil {
			return nil, fmt.Errorf("erro creating item %w", err)
		}
		questionItems[i] = item
	}

	return questionItems, nil
}

func newFromExam(ex exam.Exam) mongoExam {
	mongoQuestions := make([]*mongoQuestion, len(ex.GetQuestions()))

	for i, q := range ex.GetQuestions() {
		mongoItems := make([]*mongoQuestionItem, len(q.GetItems()))

		for j, qi := range q.GetItems() {
			mongoItems[j] = &mongoQuestionItem{
				id:    qi.GetID(),
				text:  qi.GetText(),
				right: qi.GetIsRight(),
			}
		}
		mongoQuestions[i] = &mongoQuestion{
			id:        q.GetID(),
			statement: q.GetStatement(),
			items:     mongoItems,
		}
	}

	return mongoExam{
		id:          ex.GetID(),
		description: ex.GetDescription(),
		title:       ex.GetTitle(),
		questions:   mongoQuestions,
	}

}

func New(ctx context.Context, uriConnection string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uriConnection))
	if err != nil {
		return nil, err
	}

	// TODO: Add environment vairable
	db := client.Database("quizzer")
	collectionExam := db.Collection("exam")
	repo := MongoRepository{
		db:   db,
		exam: collectionExam,
	}

	return &repo, nil
}

func (mr *MongoRepository) Get(id string) (*exam.Exam, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := mr.exam.FindOne(ctx, bson.D{{Key: "id", Value: id}})

	var me mongoExam

	if err := result.Decode(me); err != nil {
		return nil, err
	}

	return me.toAggregate()
}

func (mr *MongoRepository) GetAll() ([]*exam.Exam, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := mr.exam.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var exams []*exam.Exam
	for cursor.Next(ctx) {
		var mongoEx mongoExam

		if err := cursor.Decode(&mongoEx); err != nil {
			return nil, fmt.Errorf("erro decoding exam %w", err)
		}

		aggregateExam, err := mongoEx.toAggregate()
		if err != nil {
			return nil, err
		}

		exams = append(exams, aggregateExam)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("erro while iterating over exams %w", err)
	}

	return exams, nil
}

func (mr *MongoRepository) Save(exAdd *exam.Exam) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	doc := newFromExam(*exAdd)

	_, err := mr.exam.InsertOne(ctx, doc)
	if err != nil {
		return fmt.Errorf("erro inserting exam into mongoDB %w", err)
	}
	return nil
}

func (mr *MongoRepository) Update(updEx *exam.Exam) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updatedDoc := newFromExam(*updEx)

	_, err := mr.exam.UpdateOne(ctx, bson.D{{Key: "id", Value: updatedDoc.id}}, updatedDoc)
	if err != nil {
		return fmt.Errorf("erro updating exam %s, %w", updatedDoc.id, err)
	}

	return nil
}

func (mr *MongoRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := mr.exam.DeleteOne(ctx, bson.D{{Key: "id", Value: id}})
	if err != nil {
		return fmt.Errorf("erro deleting exam %s, %w", id, err)
	}

	return nil
}