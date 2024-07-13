package handlers

import (
	"net/http"

	"github.com/blanc42/mooca/pkg/handlers/request"
	"github.com/blanc42/mooca/pkg/handlers/response"
	"github.com/blanc42/mooca/pkg/usecases"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase usecases.UserUseCase
}

func NewUserHandler(userUseCase usecases.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (u *UserHandler) RegisterUser(c *gin.Context) {
	var user request.SingupRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.userUseCase.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	msg := "User registered successfully"

	c.JSON(http.StatusOK, response.GenericResponse{Message: &msg})
}

// login user
func (u *UserHandler) LoginUser(c *gin.Context) {
	var user request.LoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := u.userUseCase.LoginUser(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// func (u *UserHandler) GetUserProfile(c *gin.Context) {
// 	// id, _ := c.Params.Get("id")

// 	// get id from params as number
// 	id, err := strconv.Atoi(c.Params.ByName("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
// 		return
// 	}

// 	// id := c.MustGet("id").(uint)

// 	// get user from use case

// 	user, err := u.userUseCase.GetUserProfile(uint(id))
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": user})
// }
