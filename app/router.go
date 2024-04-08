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

	// allow all origins
	service.Use(func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
			return
		}
		context.Next()
	})

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(context *gin.Context) {
		context.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api/v1")
	userRouter := router.Group("/users")

	userRouter.POST("", userController.CreateUser)
	userRouter.GET("", userController.SearchUsers)
	userRouter.GET("/:id", userController.GetUser)
	userRouter.PUT("/:id", userController.UpdateUser)
	userRouter.DELETE("/:id", userController.DeleteUser)
	//userRouter.GET("/search", userController.SearchUsers)

	return service
}
