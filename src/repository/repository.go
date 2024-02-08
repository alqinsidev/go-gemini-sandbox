package repository

import (
	"database/sql"

	"github.com/alqinsidev/go-gemini-sandbox/src/domain"
	"github.com/gin-gonic/gin"
)

const (
	InformationPath = "src/pattern.json"
)

type (
	Repository struct {
		db *sql.DB
	}

	RepositoryInterface interface {
		InsertInformation(ctx *gin.Context, payload domain.InformationBodyRequest) (id int64, err error)
		EditInformation(ctx *gin.Context, payload domain.InformationBodyRequest, id int64) (isSuccess bool, err error)
		DeleteInformation(ctx *gin.Context, id int64) (success bool, err error)
		GetInformations(ctx *gin.Context) (informations []domain.Information, err error)
		GetInformationByID(ctx *gin.Context, id int64) (information domain.Information, err error)
	}
)

func Init(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
