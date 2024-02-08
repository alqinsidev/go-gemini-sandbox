package usecase

import (
	"github.com/alqinsidev/go-gemini-sandbox/src/usecase/helper"
	"github.com/gin-gonic/gin"
)

func (uc *Usecase) GetChatResponse(ctx *gin.Context, question string) (answer string, err error) {
	informations, err := uc.repo.GetInformations(ctx)
	if err != nil {
		return
	}

	model := uc.genai.GenerativeModel("gemini-pro")
	model.SetTemperature(0)
	parts := helper.PopulateParts(informations, question)

	res, err := model.GenerateContent(ctx, parts...)
	if err != nil {
		return
	}

	answer = helper.ExtractAnswer(res)

	return
}
