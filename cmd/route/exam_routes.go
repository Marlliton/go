package route

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func Exam(rg *gin.RouterGroup, controllers any) {
	group := rg.Group("exam")

	group.GET("/", func(ctx *gin.Context) {
		type ColorGroup struct {
			ID     int
			Name   string
			Colors []string
		}
		group := ColorGroup{
			ID:     1,
			Name:   "Reds",
			Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		}
		ctx.Writer.Header().Set("Content-Type", "application/json")
		b, _ := json.Marshal(group)

		ctx.Writer.Write(b)
	})
}
