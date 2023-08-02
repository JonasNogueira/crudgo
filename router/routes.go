package router

import (
	"github.com/JonasNogueira/crudgo/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/page/v1")
	{
		v1.GET("/student", handler.ShowStudentHandler)
		v1.POST("/student", handler.CreateStudentHandler)
		v1.DELETE("/student", handler.DeleteStudentHandler)
		v1.PUT("/student", handler.UpdateStudentHandler)
		v1.GET("/studentS", handler.ListStudentsHandler)

	}

}
