package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/personal/user-manager-backend/app/dto"
	"github.com/personal/user-manager-backend/app/service"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Errorf("failed to parse id: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})

		return
	}

	user, err := c.userService.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, errors.New(service.ErrUserNotFound)) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		log.Errorf("failed to fetch user: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user dto.CreateUserDto
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Errorf("failed to parse json from request body: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request: could not parse json from request body"})

		return
	}

	if user.CompanyID == nil {
		log.Error("company_id is required")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "company_id is required"})

		return
	}

	// check for user

	_, err := c.userService.FindCompanyByID(ctx, *user.CompanyID)
	if err != nil {
		log.Errorf("failed to fetch company: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to fetch company"})

		return
	}

	userWithEmail, err := c.userService.FindUserByEmail(ctx, *user.Email)
	if err != nil {

		log.Errorf("failed to fetch user by email: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user by email"})

		return
	}

	if userWithEmail != nil {
		log.Errorf("user with email %s already exists", *user.Email)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user with email already exists"})

		return
	}

	if err = c.userService.CreateUser(ctx, &user); err != nil {
		log.Errorf("failed to create user: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to create user"})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var user dto.CreateUserDto
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Error("failed to parse json from request body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request: could not parse json from request body"})

		return
	}

	// id comes from the URL
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Errorf("failed to parse id: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})

		return
	}

	user.ID = &id

	if err := c.userService.UpdateUser(ctx, &user); err != nil {
		log.Error("failed to update user")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})

		return
	}

	ctx.JSON(200, gin.H{"message": "user updated"})
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Errorf("failed to parse id: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})

		return
	}

	if err = c.userService.DeleteUser(ctx, id); err != nil {
		log.Errorf("failed to delete user: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func (c *UserController) SearchUsers(ctx *gin.Context) {
	query := ctx.Query("query")

	page, size := 1, 10

	if p, err := strconv.Atoi(ctx.Query("page")); err == nil {
		page = p
	}

	if s, err := strconv.Atoi(ctx.Query("size")); err == nil {
		size = s
	}

	users, err := c.userService.Search(ctx, query, page, size)
	if err != nil {
		log.Errorf("failed to search users: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to search users"})

		return
	}

	ctx.JSON(http.StatusOK, users)
}
