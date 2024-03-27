package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthenticateMiddleware(ctx *gin.Context) {
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

func (h *Handler) CORS(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if ctx.Request.Method == "OPTIONS" {
		ctx.Writer.WriteHeader(http.StatusOK)
		return
	}
	ctx.Next()
}
