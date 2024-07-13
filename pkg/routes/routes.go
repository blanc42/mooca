package routes

import (
	"github.com/blanc42/mooca/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(UserHandler *handlers.UserHandler, InsituteHandler *handlers.InstituteHandler, TeacherHandler *handlers.TeacherHandler, CourseHandler *handlers.CourseHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// apiRouter := r.Group("/api/v1",)
	apiRouter := r.Group("/api/v1")
	{
		apiRouter.POST("/signup", UserHandler.RegisterUser)
		apiRouter.POST("/login", UserHandler.LoginUser)

		// usersGroup := apiRouter.Group("/users")
		// {
		// 	// usersGroup.POST("", UserHandler.RegisterUser)
		// 	usersGroup.GET("/:id", UserHandler.GetUserProfile)
		// }
		// institutesGroup := apiRouter.Group("/institutes")
		// {
		// 	institutesGroup.POST("", InsituteHandler.CreateInstitute)
		// 	institutesGroup.GET("/:id", InsituteHandler.GetInstituteByID)
		// }
		// teachersGroup := apiRouter.Group("/teachers")
		// {
		// 	teachersGroup.POST("", TeacherHandler.CreateTeacher)
		// 	teachersGroup.GET("/:id", TeacherHandler.GetTeacherByID)
		// }
		// coursesGroup := apiRouter.Group("/courses")
		// {
		// 	coursesGroup.POST("", CourseHandler.CreateCourse)
		// 	coursesGroup.GET("/:id", CourseHandler.GetCourseByID)
		// }
	}

	return r
}
