package handler

import (
	"net/http"

	"github.com/JonasNogueira/crudgo/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List students
// @Description List all students
// @Tags Students
// @Accept json
// @Produce json
// @Success 200 {object} ListStudentsResponse
// @Failure 500 {object} ErrorResponse
// @Router /students [get]
func ListStudentsHandler(ctx *gin.Context) {
	students := []schemas.Student{}

	if err := db.Find(&students).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing students")
		return
	}

	sendSuccess(ctx, "list-students", students)
}
