package domain

import (
	"github.com/google/generative-ai-go/genai"
)

type GenAI struct {
	Client         *genai.Client
	PreDefinedPart []genai.Part
}

type Part struct {
	Text string `json:"text"`
}

type ChatRequestPayload struct {
	Question string `json:"question" binding:"required"`
}

type ChatResponse struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
