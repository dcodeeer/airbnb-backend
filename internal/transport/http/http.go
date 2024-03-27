package http

import (
	"api/internal/application"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const bodyMaxBody = 3 * 1024 * 1024

type Handler struct {
	useCase *application.UseCase
}

func New(useCase *application.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) Run(socket string) error {
	server := &http.Server{
		Handler:      h.router(),
		Addr:         socket,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server.ListenAndServe()
}

func (h *Handler) router() http.Handler {
	r := gin.New()

	r.Use(func(ctx *gin.Context) {
		ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, bodyMaxBody)
		ctx.Next()
	})
	r.Use(h.CORS)

	h.registerRoutes(r)

	return r
}

func (h *Handler) registerRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")

	h.setupUsers(api)
	h.setupBooking(api)
	h.setupEstates(api)
}
