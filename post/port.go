package post

import (
	"blog/domain"
	postHandler "blog/rest/handler/post"
)

type Service interface {
	postHandler.Service
}

type PostRepo interface {
	Create(p domain.Post) (*domain.Post, error)
	List(page, limit int64) ([]*domain.Post, error)
	Count() (int64, error)
	Get(id int) (*domain.Post, error)
	Update(p domain.Post) (*domain.Post, error)
	Delete(id int) error
}
