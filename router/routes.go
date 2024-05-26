package router

import (
	docs "github.com/LeandroSAlmeida/golang-api/docs"
	"github.com/LeandroSAlmeida/golang-api/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Definition BasePath
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Initialize Handler
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
