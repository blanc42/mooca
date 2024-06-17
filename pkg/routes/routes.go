package routes

import (
	"github.com/blanc42/mooca/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(UserHandler *handlers.UserHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
