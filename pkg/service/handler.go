package handler

import (
	"core-users/pkg/domain/users"
)

type Handler struct {
	userService users.Service
}

func NewHandler(userService users.Service) *Handler {
	return &Handler{userService: userService}
}
