package usecases

import (
	"github.com/blanc42/mooca/pkg/handlers/request"
	"github.com/blanc42/mooca/pkg/models"
)

type UserUseCase interface {
	RegisterUser(user *request.SingupRequest) error
	LoginUser(user *request.LoginRequest) (string, error)
	GetUserProfile(id uint) (*models.User, error)
	UpdateUserProfile(user *models.User) error
	DeleteUserAccount(id uint) error
}
