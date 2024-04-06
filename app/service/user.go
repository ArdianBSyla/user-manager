package service

import (
	"github.com/gin-gonic/gin"
	"github.com/personal/user-manager-backend/app/dto"
	log "github.com/sirupsen/logrus"

	"github.com/personal/user-manager-backend/app/repository"
)

type UserService interface {
	CreateUser(ctx *gin.Context, user *dto.CreateUserDto) error
	GetUser(ctx *gin.Context, id int) (*dto.UserDto, error)
	GetUsers(ctx *gin.Context, page, size int) ([]*dto.UserDto, error)
	UpdateUser(ctx *gin.Context, user *dto.CreateUserDto) error
	DeleteUser(ctx *gin.Context, id int) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(ctx *gin.Context, user *dto.CreateUserDto) error {
	if err := s.userRepo.Create(user); err != nil {
		log.Errorf("failed to create user: %v", err)
		return err
	}

	return nil
}

func (s *userService) GetUser(ctx *gin.Context, id int) (*dto.UserDto, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		log.Errorf("failed to find user by id: %v", err)
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUsers(ctx *gin.Context, page, size int) ([]*dto.UserDto, error) {
	users, err := s.userRepo.FindAll(page, size)
	if err != nil {
		log.Errorf("failed to fetch users: %v", err)
		return nil, err
	}

	return users, nil
}

func (s *userService) UpdateUser(ctx *gin.Context, user *dto.CreateUserDto) error {
	if err := s.userRepo.Update(user); err != nil {
		log.Errorf("failed to update user: %v", err)
		return err
	}

	return nil
}

func (s *userService) DeleteUser(ctx *gin.Context, id int) error {
	if err := s.userRepo.Delete(id); err != nil {
		log.Errorf("failed to delete user: %v", err)
		return err
	}

	return nil
}
