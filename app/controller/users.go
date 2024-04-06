package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/personal/user-manager-backend/app/dto"
	"github.com/personal/user-manager-backend/app/service"
	"strconv"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	page, size := 1, 10
	if p, err := strconv.Atoi(ctx.Query("page")); err == nil {
		page = p
	}
	if s, err := strconv.Atoi(ctx.Query("size")); err == nil {
		size = s
	}

	users, err := c.userService.GetUsers(ctx, page, size)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": users})
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	user, err := c.userService.GetUser(ctx, id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": user})
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user dto.CreateUserDto
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.userService.CreateUser(ctx, &user); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"message": "user created"})
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var user dto.CreateUserDto
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.userService.UpdateUser(ctx, &user); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "user updated"})
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	if err := c.userService.DeleteUser(ctx, id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "user deleted"})
}
