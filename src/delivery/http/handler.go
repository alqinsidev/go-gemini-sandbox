package http

import (
	"github.com/alqinsidev/go-gemini-sandbox/src/usecase"
)

type Handler struct {
	Usecase usecase.UsecaseInterface
}

func NewHandler(uc *usecase.Usecase) *Handler {
	return &Handler{
		Usecase: uc,
	}
}
