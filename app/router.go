package router

import (
	"github.com/gin-gonic/gin"
	"github.com/personal/user-manager-backend/app/controller"
	"net/http"
)

func NewRouter(
	userController *controller.UserController,
) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(context *gin.Context) {
		context.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api/v1")
	userRouter := router.Group("/users")

	userRouter.POST("", userController.CreateUser)
	userRouter.GET("", userController.GetUsers)
	userRouter.GET("/:id", userController.GetUser)
	userRouter.PUT("/:id", userController.UpdateUser)
	userRouter.DELETE("/:id", userController.DeleteUser)

	return service
}
