package dependencies

type Service interface{}

type service struct {
	repo Repository
}

func InitService(repository Repository) Service {
	return &service{
		repo: repository,
	}
}
