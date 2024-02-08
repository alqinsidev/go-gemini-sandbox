package usecase

import (
	"github.com/alqinsidev/go-gemini-sandbox/src/domain"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func (uc *Usecase) InsertInformation(ctx *gin.Context, payload domain.InformationBodyRequest) (information domain.InformationResponse, err error) {

	result, err := uc.repo.InsertInformation(ctx, payload)
	if err != nil {
		return
	}
	information = domain.InformationResponse{
		ID:       result,
		Answer:   payload.Answer,
		Question: payload.Question,
	}
	return
}

func (uc *Usecase) EditInformation(ctx *gin.Context, payload domain.InformationBodyRequest, id int64) (information domain.InformationResponse, err error) {

	result, err := uc.repo.GetInformationByID(ctx, id)
	if err != nil {
		return
	}

	_, err = uc.repo.EditInformation(ctx, payload, id)
	if err != nil {
		return
	}

	information = domain.InformationResponse{
		ID:       result.ID,
		Answer:   payload.Answer,
		Question: payload.Question,
	}

	return
}

func (uc *Usecase) DeleteInformation(ctx *gin.Context, id int64) (information domain.InformationResponse, err error) {

	result, err := uc.repo.GetInformationByID(ctx, id)
	if err != nil {
		return
	}

	_, err = uc.repo.DeleteInformation(ctx, id)
	if err != nil {
		return
	}

	information = domain.InformationResponse{
		ID:       result.ID,
		Answer:   result.Answer,
		Question: result.Question,
	}
	return
}

func (uc *Usecase) GetInformations(ctx *gin.Context) (informations []domain.InformationResponse, err error) {
	result, err := uc.repo.GetInformations(ctx)
	if err != nil {
		return
	}

	err = copier.Copy(&informations, result)
	return
}

func (uc *Usecase) GetInformationByID(ctx *gin.Context, id int64) (information domain.InformationResponse, err error) {
	result, err := uc.repo.GetInformationByID(ctx, id)
	if err != nil {
		return
	}

	err = copier.Copy(&information, result)
	return
}
