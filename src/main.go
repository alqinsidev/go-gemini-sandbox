package main

import (
	"fmt"

	"github.com/alqinsidev/go-gemini-sandbox/src/config"
	"github.com/spf13/viper"
)

func main() {
	app := config.Init()

	config.Bootstrap(&config.AppBootstrapConfig{
		Gin:   app.Gin,
		GenAI: app.GenAI,
		Ctx:   app.Ctx,
		DB:    app.DB,
	})

	appPort := fmt.Sprintf(`:%s`, viper.GetString("APP_PORT"))
	app.Gin.Run(appPort)
}
