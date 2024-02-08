package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitGin() *gin.Engine {
	var app = gin.New()
	app.Use(cors.Default())

	return app
}
