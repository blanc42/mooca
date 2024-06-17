package handlers

import (
	"net/http"

	"github.com/blanc42/mooca/pkg/models"
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
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.userUseCase.RegisterUser(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func (u *UserHandler) GetUserProfile(c *gin.Context) {
	// id, _ := c.Params.Get("id")

	// get id from params as number
	// id, err := strconv.Atoi(c.Params.ByName("id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	// 	return
	// }

	id := c.MustGet("id").(uint)

	// get user from use case

	user, err := u.userUseCase.GetUserProfile(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
