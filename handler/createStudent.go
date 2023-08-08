package handler

import (
	"net/http"

	"github.com/JonasNogueira/crudgo/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create Student
// @Description Create a new  Student
// @Tags Students
// @Accept json
// @Produce json
// @Param request body CreateStudentRequest true "Request body"
// @Success 200 {object} CreateStudentResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /student [post]
func CreateStudentHandler(ctx *gin.Context) {
	request := CreateStudentRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	student := schemas.Student{
		Name:   request.Name,
		Period: request.Period,
		Lab:    request.Lab,
	}

	if err := db.Create(&student).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", student)
}
