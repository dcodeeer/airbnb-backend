package http

import (
	"api/internal/core"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) setupUsers(r *gin.RouterGroup) {
	users := r.Group("/users")

	users.POST("/", h.usersSignUp)
	users.POST("/signin", h.usersSignIn)
	users.POST("/recovery", h.usersSendRecoveryKey)
	users.POST("/recovery/confirm", h.usersConfirmRecoveryKey)
	{
		private := users.Use(h.Authenticate)
		private.GET("/", h.usersGetMe)
		private.GET("/update", h.usersUpdate)
	}
}

func (h *Handler) usersSignUp(ctx *gin.Context) {
	var input SignUpDto
	if err := ctx.BindJSON(&input); err != nil {
		ctx.Status(400)
		return
	}

	res, err := h.useCase.Users.SignUp(ctx.Request.Context(), input.Email, input.Password)
	if err != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, gin.H{"token": res})
}

func (h *Handler) usersSignIn(ctx *gin.Context) {
	var input SignUpDto
	if err := ctx.BindJSON(&input); err != nil {
		ctx.Status(400)
		return
	}

	res, err := h.useCase.Users.SignIn(ctx.Request.Context(), input.Email, input.Password)
	if err != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, gin.H{"token": res})
}

func (h *Handler) usersSendRecoveryKey(ctx *gin.Context) {
	var input SendRecoveryDto
	if err := ctx.BindJSON(&input); err != nil {
		ctx.Status(400)
		return
	}

	if err := h.useCase.Users.SendRecoveryKey(ctx.Request.Context(), input.Email); err != nil {
		ctx.Status(400)
		log.Println(err)
		return
	}

	ctx.String(200, "password recovery link sent to your e-mail")
}

func (h *Handler) usersGetMe(ctx *gin.Context) {
	userId := h.GetUserID(ctx)

	user, err := h.useCase.Users.GetOneById(ctx.Request.Context(), userId)
	if err != nil {
		ctx.Status(404)
		return
	}

	ctx.JSON(200, user)
}

func (h *Handler) usersUpdate(ctx *gin.Context) {
	var input UpdateDto
	if err := ctx.BindJSON(&input); err != nil {
		ctx.Status(400)
		return
	}

	userId := h.GetUserID(ctx)

	user := &core.User{
		ID:         userId,
		Email:      input.Email,
		Phone:      &input.Phone,
		FirstName:  &input.FirstName,
		LastName:   &input.LastName,
		Patronymic: &input.Patronymic,
	}

	err := h.useCase.Users.Update(ctx.Request.Context(), user)
	if err != nil {
		ctx.Status(404)
		return
	}

	ctx.JSON(200, user)
}

func (h *Handler) usersConfirmRecoveryKey(ctx *gin.Context) {
	var input ConfirmRecoveryDto
	if err := ctx.BindJSON(&input); err != nil {
		ctx.Status(400)
		return
	}

	token, err := h.useCase.Users.ConfirmRecoveryKey(ctx.Request.Context(), input.Key, input.Password)
	if err != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, gin.H{"token": token})
}
