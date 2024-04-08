package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/personal/user-manager-backend/app/dto"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/personal/user-manager-backend/app/repository"
)

const ErrUserNotFound = "user not found"

type UserService interface {
	CreateUser(ctx *gin.Context, user *dto.CreateUserDto) error
	GetUser(ctx *gin.Context, id int) (*dto.UserDto, error)
	GetUsers(ctx *gin.Context, page, size int) (*dto.UserDtoPaginated, error)
	UpdateUser(ctx *gin.Context, user *dto.CreateUserDto) error
	DeleteUser(ctx *gin.Context, id int) error
	FindCompanyByID(ctx *gin.Context, id int) (*dto.CompanyDto, error)
	FindUserByEmail(ctx *gin.Context, email string) (*dto.UserDto, error)
	Search(ctx *gin.Context, query string, page, pageSize int) (*dto.UserDtoPaginated, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(ctx *gin.Context, user *dto.CreateUserDto) error {
	if err := s.userRepo.Create(ctx, user); err != nil {
		log.Errorf("failed to create user: %v", err)
		return errors.New("failed to create user")
	}

	return nil
}

func (s *userService) GetUser(ctx *gin.Context, id int) (*dto.UserDto, error) {
	log.Debug("UserService.GetUser")
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		log.Errorf("failed to find user by id: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(ErrUserNotFound)
		}

		return nil, errors.New("failed to find user by id")
	}

	return user, nil
}

func (s *userService) GetUsers(ctx *gin.Context, page, size int) (*dto.UserDtoPaginated, error) {
	users, err := s.userRepo.FindAll(ctx, page, size)
	if err != nil {
		log.Errorf("failed to fetch users: %v", err)

		return nil, errors.New("failed to fetch users")
	}

	usersCount, err := s.userRepo.Count(ctx)
	if err != nil {
		log.Errorf("failed to count users: %v", err)

		return nil, errors.New("failed to get count for users")
	}

	return &dto.UserDtoPaginated{
		Data:  users,
		Total: usersCount,
	}, nil
}

func (s *userService) UpdateUser(ctx *gin.Context, user *dto.CreateUserDto) error {
	if err := s.userRepo.Update(ctx, user); err != nil {
		log.Errorf("failed to update user: %v", err)

		return errors.New("failed to update user")
	}

	return nil
}

func (s *userService) DeleteUser(ctx *gin.Context, id int) error {
	if err := s.userRepo.Delete(ctx, id); err != nil {
		log.Errorf("failed to delete user: %v", err)

		return err
	}

	return nil
}

func (s *userService) FindCompanyByID(ctx *gin.Context, id int) (*dto.CompanyDto, error) {
	company, err := s.userRepo.FindCompanyByID(ctx, id)
	if err != nil {
		log.Errorf("failed to find company by id: %v", err)

		return nil, errors.New("failed to find company by id")
	}

	return company, nil
}

func (s *userService) FindUserByEmail(ctx *gin.Context, email string) (*dto.UserDto, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		log.Errorf("failed to find user by email: %v", err)

		return nil, errors.New("failed to find user by email")
	}

	return user, nil
}

func (s *userService) Search(ctx *gin.Context, query string, page, pageSize int) (*dto.UserDtoPaginated, error) {
	users, count, err := s.userRepo.SearchAndCount(ctx, page, pageSize, query)
	if err != nil {
		log.Errorf("failed to search users: %v", err)

		return nil, errors.New("failed to search users")
	}

	return &dto.UserDtoPaginated{
		Data:  users,
		Total: count,
	}, nil
}
