package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmazleo/opportunities/schemas"
)

// @Summary Mostra uma abertura
// @Description Retorna detalhes de uma abertura
// @Tags Opening
// @Accept json
// @Produce json
// @Success 200 {object} handler.Opening
// @Router /opening [get]

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if  id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error;err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
