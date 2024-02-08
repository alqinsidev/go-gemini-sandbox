package http

import (
	"net/http"

	"github.com/alqinsidev/go-gemini-sandbox/src/domain"
	_error "github.com/alqinsidev/go-gemini-sandbox/src/error"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetChatResponse(ctx *gin.Context) {
	var body domain.ChatRequestPayload
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": _error.ErrBadPayloadRequest.Error(),
		})
		return
	}

	answer, err := h.Usecase.GetChatResponse(ctx, body.Question)
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
			Question: body.Question,
			Answer:   answer,
		},
	}
	ctx.JSON(http.StatusOK, response)
}
