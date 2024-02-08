package route

import (
	"github.com/alqinsidev/go-gemini-sandbox/src/delivery/http"
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	Gin     *gin.Engine
	Handler *http.Handler
}

func (c *RouteConfig) Setup() {
	c.Gin.GET("/informations", c.Handler.GetInformation)
	c.Gin.GET("/informations/:id", c.Handler.GetInformationByID)
	c.Gin.POST("/informations", c.Handler.InsertInformation)
	c.Gin.PUT("/informations/:id", c.Handler.EditInformation)
	c.Gin.DELETE("/informations/:id", c.Handler.DeleteInformation)

	c.Gin.POST("/chat", c.Handler.GetChatResponse)
}
