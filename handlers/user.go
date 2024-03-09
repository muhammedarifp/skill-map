package handlers

import "github.com/muhammedarifp/skill-map/managers"

type UserHandler struct {
	userManager managers.UserManager
}

func NewUserHandlerFrom(um managers.UserManager) *UserHandler {
	return &UserHandler{
		userManager: um,
	}
}
