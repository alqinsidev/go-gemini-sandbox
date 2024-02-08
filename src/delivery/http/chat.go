package http

import (
	"net/http"

	"github.com/alqinsidev/go-gemini-sandbox/src/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetChatResponse(ctx *gin.Context) {
	var body domain.ChatRequestPayload
	ctx.Bind(&body)

	answer, err := h.Usecase.GetChatResponse(ctx, body.Input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	response := gin.H{
		"success": true,
		"data": domain.ChatResponse{
			Input:  body.Input,
			Output: answer,
		},
	}
	ctx.JSON(http.StatusOK, response)
}
