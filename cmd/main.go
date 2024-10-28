package main

import (
	"github.com/Marlliton/go-quizzer/cmd/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	route.Exam(&r.RouterGroup, "")

	r.Run(":8000")

}
