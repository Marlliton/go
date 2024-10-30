package main

import (
	"context"
	"time"

	"github.com/Marlliton/go-quizzer/infra/api/controller"
	"github.com/Marlliton/go-quizzer/infra/api/route"
	"github.com/Marlliton/go-quizzer/services/examsvc"
	"github.com/gin-gonic/gin"
)

var (
	ctx           context.Context
	uriConnection = "mongodb://localhost:27017"
)

func main() {
	r := gin.Default()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	examServ, err := examsvc.NewExamsvc(
		examsvc.WithMongoExamRepository(ctx, uriConnection),
	)
	if err != nil {
		panic(err)
	}

	route.Exam(&r.RouterGroup, controller.NewExamController(*examServ))

	r.Run(":8000")

}
