package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/personal/user-manager-backend/app/model"
	"gorm.io/gorm"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/personal/user-manager-backend/app/dto"
)

type UserRepository interface {
	Create(ctx *gin.Context, user *dto.CreateUserDto) error
	Update(ctx *gin.Context, user *dto.CreateUserDto) error
	FindByID(ctx *gin.Context, id int) (*dto.UserDto, error)
	Count(ctx *gin.Context) (int, error)
	FindAll(ctx *gin.Context, page, size int) ([]*dto.UserDto, error)
	Delete(ctx *gin.Context, id int) error
	FindCompanyByID(ctx *gin.Context, id int) (*dto.CompanyDto, error)
	FindByEmail(ctx *gin.Context, email string) (*dto.UserDto, error)
	SearchAndCount(ctx *gin.Context, page, pagesize int, query string) ([]*dto.UserDto, int, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx *gin.Context, user *dto.CreateUserDto) error {
	email := strings.ToLower(*user.Email)

	if err := r.db.WithContext(ctx).Model(model.User{}).Save(&model.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     &email,
		CompanyID: *user.CompanyID,
	}).Error; err != nil {
		log.Errorf("failed to create user: %v", err)

		return err
	}

	return nil
}

func (r *userRepository) Update(ctx *gin.Context, user *dto.CreateUserDto) error {
	email := strings.ToLower(*user.Email)
	user.Email = &email

	if err := r.db.WithContext(ctx).Model(model.User{}).Where("id = ?", *user.ID).Updates(model.User{
		ID:        *user.ID,
		CompanyID: *user.CompanyID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}).Error; err != nil {
		log.Errorf("failed to update user: %v", err)

		return err
	}

	return nil
}

func (r *userRepository) FindByID(ctx *gin.Context, id int) (*dto.UserDto, error) {
	var user model.User

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		log.Errorf("failed to find user by id: %v", err)
		return nil, err
	}

	return &dto.UserDto{
		ID:        user.ID,
		CompanyID: user.CompanyID,
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		Email:     *user.Email,
	}, nil
}

func (r *userRepository) FindAll(ctx *gin.Context, page, size int) ([]*dto.UserDto, error) {
	var users []*model.User

	if err := r.db.WithContext(ctx).Offset((page - 1) * size).Limit(size).Find(&users).Error; err != nil {
		log.Errorf("failed to find all users: %v", err)

		return nil, err
	}

	var dtos []*dto.UserDto
	for _, user := range users {
		dtos = append(dtos, &dto.UserDto{
			ID:        user.ID,
			CompanyID: user.CompanyID,
			FirstName: *user.FirstName,
			LastName:  *user.LastName,
			Email:     *user.Email,
		})
	}

	return dtos, nil
}

func (r *userRepository) Delete(ctx *gin.Context, id int) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		log.Errorf("failed to delete user: %v", err)

		return err
	}

	return nil
}

func (r *userRepository) Count(ctx *gin.Context) (int, error) {
	var usersCount int64
	if err := r.db.WithContext(ctx).Model(&model.User{}).Count(&usersCount).Error; err != nil {
		log.Errorf("failed to count all users: %v", err)

		return 0, err
	}

	return int(usersCount), nil
}

func (r *userRepository) FindCompanyByID(ctx *gin.Context, id int) (*dto.CompanyDto, error) {
	var company model.Company

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&company).Error; err != nil {
		log.Errorf("failed to find company by id: %v", err)

		return nil, err
	}

	return &dto.CompanyDto{
		ID: company.ID,
	}, nil
}

func (r *userRepository) FindByEmail(ctx *gin.Context, email string) (*dto.UserDto, error) {
	var user model.User

	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &dto.UserDto{
		ID:        user.ID,
		CompanyID: user.CompanyID,
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		Email:     *user.Email,
	}, nil
}

func (r *userRepository) SearchAndCount(ctx *gin.Context, page, pagesize int, query string) ([]*dto.UserDto, int, error) {
	var users []*model.User

	if err := r.db.WithContext(ctx).Where("first_name LIKE ? OR last_name LIKE ? OR email LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%").Offset((page - 1) * pagesize).Limit(pagesize).Find(&users).Error; err != nil {
		log.Errorf("failed to search users: %v", err)

		return nil, 0, err
	}

	var dtos []*dto.UserDto
	for _, user := range users {
		dtos = append(dtos, &dto.UserDto{
			ID:        user.ID,
			CompanyID: user.CompanyID,
			FirstName: *user.FirstName,
			LastName:  *user.LastName,
			Email:     *user.Email,
		})
	}

	var count int64
	if err := r.db.WithContext(ctx).Model(&model.User{}).Where("first_name LIKE ? OR last_name LIKE ? OR email LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%").Count(&count).Error; err != nil {
		log.Errorf("failed to count search results: %v", err)

		return nil, 0, err
	}

	return dtos, int(count), nil
}
