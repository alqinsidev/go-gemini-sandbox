package usecase

import (
	"github.com/alqinsidev/go-gemini-sandbox/src/domain"
	"github.com/alqinsidev/go-gemini-sandbox/src/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
)

type (
	Usecase struct {
		repo  repository.RepositoryInterface
		genai *genai.Client
	}

	UsecaseInterface interface {
		InsertInformation(ctx *gin.Context, payload domain.InformationBodyRequest) (information domain.InformationResponse, err error)
		GetInformations(ctx *gin.Context) (informations []domain.InformationResponse, err error)
		GetInformationByID(ctx *gin.Context, id int64) (information domain.InformationResponse, err error)
		EditInformation(ctx *gin.Context, payload domain.InformationBodyRequest, id int64) (information domain.InformationResponse, err error)
		DeleteInformation(ctx *gin.Context, id int64) (information domain.InformationResponse, err error)

		GetChatResponse(ctx *gin.Context, question string) (answer string, err error)
	}
)

func Init(genai *genai.Client, repo repository.RepositoryInterface) *Usecase {
	return &Usecase{
		repo:  repo,
		genai: genai,
	}
}
