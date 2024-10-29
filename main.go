package main

import (
	"github.com/Marlliton/go-quizzer/infra/api/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	route.Exam(&r.RouterGroup, "")

	r.Run(":8000")

}
