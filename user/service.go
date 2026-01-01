package user

import "blog/domain"

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (svc *service) Create(u domain.User) (*domain.User, error) {
	user, err := svc.userRepo.Create(u)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return user, nil
}

func (svc *service) Find(email, password string) (*domain.User, error) {
	user, err := svc.userRepo.Find(email, password)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return user, nil
}
