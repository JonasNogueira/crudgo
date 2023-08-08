package handler

import (
	"net/http"

	"github.com/JonasNogueira/crudgo/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show Student
// @Description Show Student
// @Tags Students
// @Accept json
// @Produce json
// @Param id query string true "Student identification"
// @Success 200 {object} ShowStudentResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /student [get]
func ShowStudentHandler(ctx *gin.Context) {
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

	sendSuccess(ctx, "show-student", student)
}
