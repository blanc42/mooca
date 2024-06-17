package usecases

import (
	"github.com/blanc42/mooca/pkg/models"
	"github.com/blanc42/mooca/pkg/repository"
)

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{userRepository: userRepository}
}

func (u *userUseCase) RegisterUser(user *models.User) error {
	return u.userRepository.CreateUser(user)
}

func (u *userUseCase) GetUserProfile(id uint) (*models.User, error) {
	return u.userRepository.GetUserByID(id)
}

func (u *userUseCase) UpdateUserProfile(user *models.User) error {
	return u.userRepository.UpdateUser(user)
}

func (u *userUseCase) DeleteUserAccount(id uint) error {
	return u.userRepository.DeleteUser(id)
}
