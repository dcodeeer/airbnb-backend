package http

import "github.com/gin-gonic/gin"

func (h *Handler) setupEstates(r *gin.RouterGroup) {
	users := r.Group("/estates")
	users.GET("", h.estatesGet)
}

func (h *Handler) estatesGet(ctx *gin.Context) {
	estates, err := h.useCase.Estates.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.Status(500)
		return
	}

	ctx.JSON(200, estates)
}
