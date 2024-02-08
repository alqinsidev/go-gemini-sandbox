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
	Input string `json:"input"`
}

type ChatResponse struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}
