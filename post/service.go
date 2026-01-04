package post

import "blog/domain"

type service struct {
	postRepo PostRepo
}

func NewService(postRepo PostRepo) Service {
	return &service{
		postRepo: postRepo,
	}
}

func (svc service) Create(p domain.Post) (*domain.Post, error) {
	return svc.postRepo.Create(p)
}

func (svc service) List() ([]*domain.Post, error) {
	return svc.postRepo.List()
}

func (svc service) Get(id int) (*domain.Post, error) {
	return svc.postRepo.Get(id)
}

func (svc service) Update(p domain.Post) (*domain.Post, error) {
	return svc.postRepo.Update(p)
}

func (svc service) Delete(id int) error {
	return svc.postRepo.Delete(id)
}
