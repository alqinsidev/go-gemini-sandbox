package http

import (
	"net/http"
	"strconv"

	"github.com/alqinsidev/go-gemini-sandbox/src/domain"
	_error "github.com/alqinsidev/go-gemini-sandbox/src/error"
	"github.com/gin-gonic/gin"
)

func (h *Handler) InsertInformation(ctx *gin.Context) {
	var body domain.InformationBodyRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": _error.ErrBadPayloadRequest.Error(),
		})
		return
	}

	information, err := h.Usecase.InsertInformation(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"success": true,
		"data":    information,
	}
	ctx.JSON(http.StatusCreated, response)
}

func (h *Handler) EditInformation(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": _error.ErrBadPayloadRequest.Error(),
		})
		return
	}

	var body domain.InformationBodyRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": _error.ErrBadPayloadRequest.Error(),
		})
		return
	}

	information, err := h.Usecase.EditInformation(ctx, body, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"success": true,
		"data":    information,
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteInformation(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": _error.ErrBadPayloadRequest.Error(),
		})
		return
	}

	information, err := h.Usecase.DeleteInformation(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"success": true,
		"data":    information,
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetInformationByID(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": _error.ErrBadPayloadRequest.Error(),
		})
		return
	}

	information, err := h.Usecase.GetInformationByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"success": true,
		"data":    information,
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetInformation(ctx *gin.Context) {
	informations, err := h.Usecase.GetInformations(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"success": true,
		"data":    informations,
	}
	ctx.JSON(http.StatusOK, response)
}
