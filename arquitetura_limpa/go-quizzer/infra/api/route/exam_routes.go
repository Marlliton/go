package route

import (
	"github.com/Marlliton/go-quizzer/infra/api/controller"
	"github.com/gin-gonic/gin"
)

func Exam(rg *gin.RouterGroup, controllers controller.ExamController) {
	group := rg.Group("exam")

	group.GET("/:id", controllers.Get)
	group.GET("/", controllers.GetAll)
	group.POST("/", controllers.Save)
	group.PUT("/", controllers.Update)
	group.DELETE("/", controllers.Delete)
}
