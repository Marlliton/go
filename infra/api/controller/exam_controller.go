package controller

import (
	"net/http"

	"github.com/Marlliton/go-quizzer/infra/api/dto"
	"github.com/Marlliton/go-quizzer/infra/api/httperror"
	"github.com/Marlliton/go-quizzer/services/examsvc"
	"github.com/gin-gonic/gin"
)

type (
	ExamController interface {
		Get(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		Save(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	examController struct {
		examSvc examsvc.ExamService
	}
)

func NewExamController(svc examsvc.ExamService) ExamController {
	return &examController{
		examSvc: svc,
	}
}

func (ec *examController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Id is required"})
		return
	}
	exam, err := ec.examSvc.Get(id)
	if err != nil {
		httperror.WriteError(err, ctx.Writer)
		return
	}

	response := dto.ToExamDTOResponse(*exam)

	ctx.JSON(http.StatusOK, response)
}
func (ec *examController) GetAll(ctx *gin.Context) {

}
func (ec *examController) Save(ctx *gin.Context) {

}
func (ec *examController) Update(ctx *gin.Context) {

}
func (ec *examController) Delete(ctx *gin.Context) {

}
