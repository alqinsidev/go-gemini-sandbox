package config

import (
	"context"
	"database/sql"

	"github.com/alqinsidev/go-gemini-sandbox/src/delivery/http"
	"github.com/alqinsidev/go-gemini-sandbox/src/delivery/http/route"
	"github.com/alqinsidev/go-gemini-sandbox/src/repository"
	"github.com/alqinsidev/go-gemini-sandbox/src/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/viper"
)

type AppBootstrapConfig struct {
	Gin       *gin.Engine
	Viper     *viper.Viper
	GenAI     *genai.Client
	Ctx       context.Context
	DB        *sql.DB
	Validator *validator.Validate
}

func Init() *AppBootstrapConfig {
	ctx := context.Background()
	viper := InitViper()
	gin := InitGin()
	genAI := InitGenAIClient(ctx, viper)
	DB := InitMySQL(viper)

	return &AppBootstrapConfig{
		Ctx:   ctx,
		Gin:   gin,
		GenAI: genAI,
		Viper: viper,
		DB:    DB,
	}

}

func Bootstrap(config *AppBootstrapConfig) {
	Repository := repository.Init(config.DB)
	Usecase := usecase.Init(config.GenAI, Repository)
	Handler := http.NewHandler(Usecase)

	routeConfig := route.RouteConfig{
		Gin:     config.Gin,
		Handler: Handler,
	}

	routeConfig.Setup()
}

func (app *AppBootstrapConfig) GetGin() *gin.Engine {
	return app.Gin
}

func (app *AppBootstrapConfig) GetContext() context.Context {
	return app.Ctx
}
