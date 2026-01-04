package post

import "blog/domain"

type Service interface {
	Create(p domain.Post) (*domain.Post, error)
	List() ([]*domain.Post, error)
	Get(id int) (*domain.Post, error)
	Update(p domain.Post) (*domain.Post, error)
	Delete(id int) error
}
