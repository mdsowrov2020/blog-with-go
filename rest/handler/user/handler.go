package user

import (
	"blog/config"
)

type Handler struct {
	svc Service
	cnf *config.Config
}

func NewHandler(
	svc Service,
	cnf *config.Config,
) *Handler {
	return &Handler{
		svc: svc,
		cnf: cnf,
	}
}
