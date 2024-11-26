package controller

import (
	"net/http"

	"github.com/Marlliton/go-quizzer/infra/api/dto"
	"github.com/Marlliton/go-quizzer/infra/api/httperror"
	"github.com/Marlliton/go-quizzer/infra/api/mapper"
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

	ctx.JSON(http.StatusOK, mapper.ToExamDTO(*exam))
}
func (ec *examController) GetAll(ctx *gin.Context) {
	exams, err := ec.examSvc.GetAll()
	if err != nil {
		httperror.WriteError(err, ctx.Writer)
		return
	}

	result := make([]*dto.ExamDTOResponse, 0)
	for _, ex := range exams {
		result = append(result, mapper.ToExamDTO(*ex))
	}
	ctx.JSON(http.StatusOK, result)

}
func (ec *examController) Save(ctx *gin.Context) {
	var input dto.ExamDTORequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exam, err := mapper.ToExamDomain(&input)
	if err != nil {
		httperror.WriteError(err, ctx.Writer)
		return
	}

	err = ec.examSvc.Save(exam)
	if err != nil {
		httperror.WriteError(err, ctx.Writer)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
func (ec *examController) Update(ctx *gin.Context) {

}
func (ec *examController) Delete(ctx *gin.Context) {

}
