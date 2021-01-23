package controllers

import (
	"github.com/joaopandolfi/go_generic_api/dao"
	"github.com/joaopandolfi/go_generic_api/services"
)

// NewUserService - Factory
func NewUserService() services.UserService {
	return services.User{
		UserDAO: dao.User{},
	}
}
