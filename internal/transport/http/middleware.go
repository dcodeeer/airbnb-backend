package http

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) Authenticate(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")

	user, err := h.useCase.Users.GetByToken(ctx, token)
	if err != nil {
		ctx.AbortWithStatus(401)
		return
	}

	ctx.Set("userId", user.ID)

	ctx.Next()
}

func (h *Handler) GetUserID(ctx *gin.Context) int {
	value, err := ctx.Get("userId")
	if !err {
		return 0
	}

	return value.(int)
}
