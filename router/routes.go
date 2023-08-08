package router

import (
	docs "github.com/JonasNogueira/crudgo/docs"
	"github.com/JonasNogueira/crudgo/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group("/page/v1")
	{
		v1.GET("/student", handler.ShowStudentHandler)
		v1.POST("/student", handler.CreateStudentHandler)
		v1.DELETE("/student", handler.DeleteStudentHandler)
		v1.PUT("/student", handler.UpdateStudentHandler)
		v1.GET("/studentS", handler.ListStudentsHandler)

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
