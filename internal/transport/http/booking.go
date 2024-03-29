package http

import "github.com/gin-gonic/gin"

func (h *Handler) setupBooking(r *gin.RouterGroup) {
	users := r.Group("/booking")
	users.POST("", h.bookingCreateEstate)
}

func (h *Handler) bookingCreateEstate(ctx *gin.Context) {
	ctx.String(200, "created")
}
