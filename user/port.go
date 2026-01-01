package user

import (
	"blog/domain"
	userHandler "blog/rest/handler/user"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(p domain.User) (*domain.User, error)
	Find(email, password string) (*domain.User, error)
}
