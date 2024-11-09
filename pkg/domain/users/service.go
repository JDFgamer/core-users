package users

import (
	"context"
	"core-users/pkg/domain/repository"
)

type Service interface {
	GetAllUser(ctx context.Context) error
}

type service struct {
	userRepo repository.User
}

func NewUserService(userRepo repository.User) Service {
	return &service{userRepo: userRepo}
}
