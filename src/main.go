package main

import (
	"fmt"

	"github.com/alqinsidev/go-gemini-sandbox/src/config"
)

func main() {
	app := config.Init()

	config.Bootstrap(&config.AppBootstrapConfig{
		Gin:   app.Gin,
		GenAI: app.GenAI,
		Ctx:   app.Ctx,
		DB:    app.DB,
	})

	appPort := fmt.Sprintf(`:%s`, app.Viper.GetString("APP_PORT"))
	app.Gin.Run(appPort)
}
