package repository

import "context"

type User interface {
	GetAllUser(ctx context.Context) error
}
