package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStudentHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST Student",
	})
}
