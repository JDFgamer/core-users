package dependencies

type Repository interface{}

type repository struct {
	ext External
}

func InitRepository(external External) Repository {
	return &repository{
		ext: external,
	}
}
