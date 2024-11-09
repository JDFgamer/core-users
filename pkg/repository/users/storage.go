package users

import "go.mongodb.org/mongo-driver/mongo"

type Storage struct {
	storage mongo.Database
}

func NewStorage(storage mongo.Database) Storage {
	return Storage{
		storage: storage,
	}
}
