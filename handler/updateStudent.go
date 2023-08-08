package handler

import (
	"net/http"

	"github.com/JonasNogueira/crudgo/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update Student
// @Description Update a job Student
// @Tags Students
// @Accept json
// @Produce json
// @Param id query string true "Student Identification"
// @Param Student body UpdateStudentRequest true "Student data to Update"
// @Success 200 {object} UpdateStudentResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /student [put]
func UpdateStudentHandler(ctx *gin.Context) {
	request := UpdateStudentRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	student := schemas.Student{}

	if err := db.First(&student, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "student not found")
		return
	}
	// Update opening
	if request.Name != "" {
		student.Name = request.Name
	}

	if request.Period != "" {
		student.Period = request.Period
	}

	if request.Lab != "" {
		student.Lab = request.Lab
	}

	// Save opening
	if err := db.Save(&student).Error; err != nil {
		logger.Errorf("error updating student: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating student")
		return
	}
	sendSuccess(ctx, "update-student", student)
}
