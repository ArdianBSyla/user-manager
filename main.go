package main

import (
	router "github.com/personal/user-manager-backend/app"
	"github.com/personal/user-manager-backend/app/controller"
	"github.com/personal/user-manager-backend/app/helper"
	"github.com/personal/user-manager-backend/app/repository"
	"github.com/personal/user-manager-backend/app/service"
	"net/http"
	"time"
)

func main() {

	// Database connection
	db := helper.NewGormDB()

	// Initialize the repository
	userRepo := repository.NewUserRepository(db)

	// Initialize the service
	userService := service.NewUserService(userRepo)

	// Initialize the controller
	userController := controller.NewUserController(userService)

	routes := router.NewRouter(userController)

	server := &http.Server{
		Addr:           ":3000",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
