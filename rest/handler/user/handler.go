package user

import (
	"blog/config"
	"blog/repo"
)

type Handler struct {
	userRepo repo.UserRepo
	cnf      *config.Config
}

func NewHandler(
	userRepo repo.UserRepo,
	cnf *config.Config,
) *Handler {
	return &Handler{
		userRepo: userRepo,
		cnf:      cnf,
	}
}
