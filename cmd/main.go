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
	// Routes
	r := routes.SetupRouter(userHandler, instituteHandler)

	r.Run(":8000")
}
