package service

import (
	"slate/internal/domain"
	"slate/internal/repo"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService(userRepo *repo.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) GetUserByUsernameService(username string) (*domain.User, error) {
	return u.userRepo.GetUserByUsername(username)
}
