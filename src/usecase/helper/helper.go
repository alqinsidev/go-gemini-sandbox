package helper

import (
	"fmt"

	"github.com/alqinsidev/go-gemini-sandbox/src/domain"
	"github.com/google/generative-ai-go/genai"
)

func PopulateParts(informations []domain.Information, question string) []genai.Part {
	populatedParts := make([]genai.Part, 0)
	for _, information := range informations {
		inputText := fmt.Sprintf(`input: %s`, information.Question)
		populatedParts = append(populatedParts, genai.Text(inputText))

		outputText := fmt.Sprintf(`output: %s`, information.Answer)
		populatedParts = append(populatedParts, genai.Text(outputText))
	}

	inputText := fmt.Sprintf(`input: %s`, question)
	populatedParts = append(populatedParts, genai.Text(inputText))

	populatedParts = append(populatedParts, genai.Text("output: "))

	return populatedParts
}

func ExtractAnswer(res *genai.GenerateContentResponse) string {
	for _, candidate := range res.Candidates {
		answers := candidate.Content.Parts
		if len(answers) != 0 {
			return fmt.Sprintf(`%v`, answers[0])
		}
	}
	return "tidak ada jawaban"
}
