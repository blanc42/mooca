package usecases

import "github.com/blanc42/mooca/pkg/models"

type UserUseCase interface {
	RegisterUser(user *models.User) error
	GetUserProfile(id uint) (*models.User, error)
	UpdateUserProfile(user *models.User) error
	DeleteUserAccount(id uint) error
}
