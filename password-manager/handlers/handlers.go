package handlers

import (
	"log"
	"password-manager/services"
)

type Handler struct {
	Logger        *log.Logger
	MasterService services.MasterService
	UserService   services.UserService
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{Logger: logger}
}

func (h *Handler) MasterHandler(ms services.MasterService) *Handler {
	h.MasterService = ms
	return h
}

func (h *Handler) UserHandler(us services.UserService) *Handler {
	h.UserService = us
	return h
}
