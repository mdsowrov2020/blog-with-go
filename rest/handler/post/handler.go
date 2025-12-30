package post

import (
	"blog/repo"
	"blog/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	postRepo    repo.PostRepo
}

func NewHandler(
	postRepo repo.PostRepo,
	middlewares *middleware.Middlewares,
) *Handler {
	return &Handler{
		postRepo:    postRepo,
		middlewares: middlewares,
	}
}
