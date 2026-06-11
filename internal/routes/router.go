package routes

import (
	"encurtador-url-go/internal/handler"

	"github.com/gin-gonic/gin"
)

func ConfigRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", handler.CheckHealth)
	router.POST("/gerar-codigo", handler.EncurtarUrlHandler)
	router.GET("/:codigo", handler.RetornarURLPorCodigoHandler)
	router.GET("/teste", handler.EndpointPesado)

	return router
}
