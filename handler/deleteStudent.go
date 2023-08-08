package handler

import (
	"fmt"
	"net/http"

	"github.com/JonasNogueira/crudgo/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete student
// @Description Delete a new  Student
// @Tags Students
// @Accept json
// @Produce json
// @Param id query string true "Student identification"
// @Success 200 {object} DeleteStudentResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /student [delete]
func DeleteStudentHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	student := schemas.Student{}
	// Find student
	if err := db.First(&student, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("student with id: %s not found", id))
		return
	}
	// Delete student
	if err := db.Delete(&student).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting student with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-student", student)
}
