package main

import (
	router "github.com/personal/user-manager-backend/app"
	"github.com/personal/user-manager-backend/app/controller"
	"github.com/personal/user-manager-backend/app/helper"
	"github.com/personal/user-manager-backend/app/repository"
	"github.com/personal/user-manager-backend/app/service"
	"net/http"
	"os"
	"time"
)

func main() {
	// Getting the db name field from env file
	dbName := os.Getenv("DB_NAME")         // user_manager
	dbUser := os.Getenv("DB_USER")         // root
	dbPassword := os.Getenv("DB_PASSWORD") // password
	dbHost := os.Getenv("DB_HOST")         // localhost
	dbPort := os.Getenv("DB_PORT")         // 3306

	// Getting the db config from env file
	dbConfig := helper.DBConfig{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		DBName:   dbName,
	}

	// Database connection
	db := helper.NewGormDB(dbConfig)

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
