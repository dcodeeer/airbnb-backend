package http

import (
	"api/internal/application"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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

	h.registerRoutes(r)

	return r
}

func (h *Handler) registerRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")

	h.setupUsers(api)
	h.setupBooking(api)
}
