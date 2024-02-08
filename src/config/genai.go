package config

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

func InitGenAIClient(ctx context.Context, viper *viper.Viper) *genai.Client {
	apiKey := viper.GetString("API_KEY")

	client, err := genai.NewClient(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	return client
}
