package dependencies

import (
	"core-users/pkg/repository/users"
)

type Repository struct {
	userRepo users.Storage
}

func InitRepository(external External) Repository {
	return Repository{
		userRepo: users.NewStorage(*external.mongo),
	}
}
