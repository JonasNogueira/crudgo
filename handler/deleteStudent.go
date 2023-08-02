package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteStudentHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Student",
	})
}
