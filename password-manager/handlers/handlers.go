package handlers

import (
	"log"
	"password-manager/services"
)

type Handler struct {
	Logger        *log.Logger
	MasterService services.MasterService
	UserService   services.UserService
	AdminService  services.AdminService
	LoginService  services.LoginService
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{Logger: logger}
}

func (h *Handler) AdminHandler(ad services.AdminService) *Handler {
	h.AdminService = ad
	return h
}

func (h *Handler) LoginHandler(lg services.LoginService) *Handler {
	h.LoginService = lg
	return h
}

func (h *Handler) MasterHandler(ms services.MasterService) *Handler {
	h.MasterService = ms
	return h
}

func (h *Handler) UserHandler(us services.UserService) *Handler {
	h.UserService = us
	return h
}
