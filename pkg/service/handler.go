package handler

import (
	"core-users/pkg/dependencies"
)

type Handler struct {
	service dependencies.Service
}

func NewHandler(service dependencies.Service) *Handler {
	return &Handler{
		service: service,
	}
}
