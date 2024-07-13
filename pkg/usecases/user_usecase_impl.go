package usecases

import (
	"github.com/blanc42/mooca/pkg/handlers/request"
	"github.com/blanc42/mooca/pkg/models"
	"github.com/blanc42/mooca/pkg/repository"
)

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{userRepository: userRepository}
}

func (u *userUseCase) RegisterUser(signup_request *request.SingupRequest) error {
	var user models.User
	user.Username = signup_request.Username
	user.Password = signup_request.Password
	user.Email = signup_request.Email

	user.Role = models.Role{
		Id:   2,
		Name: "user",
	}

	return u.userRepository.CreateUser(&user)
}

func (u *userUseCase) LoginUser(login_request *request.LoginRequest) (string, error) {

	// convert DTO to model
	// user := models.User{
	// 	Username: login_request.Username,
	// 	Password: login_request.Password,
	// }

	// check if the the user exists

	// if not then status unauthorized

	// if user exists then compare the hash of the password

	// if password is not correct then status unauthorized

	// if password is correct then generate a jwt token

	return "token", nil

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
