package post

import (
	"blog/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	svc         Service
}

func NewHandler(
	svc Service,
	middlewares *middleware.Middlewares,
) *Handler {
	return &Handler{
		svc:         svc,
		middlewares: middlewares,
	}
}
