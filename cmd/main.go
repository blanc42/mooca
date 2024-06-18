package main

import (
	"github.com/blanc42/mooca/pkg/db"
	"github.com/blanc42/mooca/pkg/handlers"
	"github.com/blanc42/mooca/pkg/repository"
	"github.com/blanc42/mooca/pkg/routes"
	"github.com/blanc42/mooca/pkg/usecases"
)

func init() {
	db.ConnectToDB()
	db.InitializeDB()
}

func main() {

	userRepo := repository.NewUserRepository(db.DB)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(userUseCase)

	instituteRepo := repository.NewInstituteRepository(db.DB)
	instituteUseCase := usecases.NewInstituteUseCase(instituteRepo)
	instituteHandler := handlers.NewInstituteHandler(instituteUseCase)

	teacherRepo := repository.NewTeacherRepository(db.DB)
	teacherUseCase := usecases.NewTeacherUsecase(teacherRepo)
	teacherHandler := handlers.NewTeacherHandler(teacherUseCase)

	courseRepo := repository.NewCourseRepository(db.DB)
	courseUseCase := usecases.NewCourseUsecase(courseRepo)
	courseHandler := handlers.NewCourseHandler(courseUseCase)

	// Routes
	r := routes.SetupRouter(userHandler, instituteHandler, teacherHandler, courseHandler)

	r.Run(":8000")
}
