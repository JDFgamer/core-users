package dependencies

import (
	"core-users/pkg/domain/users"
)

type Domains struct {
	UserService users.Service
}

func InitService(repository Repository) Domains {
	userService := users.NewUserService(repository.userRepo)
	return Domains{
		UserService: userService,
	}
}
