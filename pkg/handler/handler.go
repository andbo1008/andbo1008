package handler

import "testbank/pkg/service"

type (
	Handler struct {
		service service.Service
	}
)

func NewHandler(service service.Service) Handler {
	return Handler{service: service}
}
