package router

import (
	docs "github.com/LeandroSAlmeida/golang-api/docs"
	"github.com/LeandroSAlmeida/golang-api/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Definição do BasePath
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Inicialização do Handler
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		// Certifique-se de que todos os manipuladores de rota tenham anotações Swagger apropriadas
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}

	// Inicialização do Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Se houver algum erro ao configurar o Swagger, exiba uma mensagem de erro
	if err := setupSwagger(); err != nil {
		panic("Failed to set up Swagger: " + err.Error())
	}
}

// setupSwagger configura o Swagger
func setupSwagger() error {
	// Configurações do SwaggerInfo
	docs.SwaggerInfo.Title = "Golang API"
	docs.SwaggerInfo.Description = "API Documentation for Golang API"
	docs.SwaggerInfo.Version = "1.0"
	// Certifique-se de que a configuração do BasePath está correta
	// docs.SwaggerInfo.BasePath = "/api/v1"
	// Defina outros campos conforme necessário

	return nil
}
