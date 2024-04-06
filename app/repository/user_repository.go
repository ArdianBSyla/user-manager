package repository

import (
	"github.com/personal/user-manager-backend/app/model"
	"gorm.io/gorm"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/personal/user-manager-backend/app/dto"
)

type UserRepository interface {
	Create(user *dto.CreateUserDto) error
	Update(user *dto.CreateUserDto) error
	FindByID(id int) (*dto.UserDto, error)
	FindAll(page, size int) ([]*dto.UserDto, error)
	Delete(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *dto.CreateUserDto) error {
	email := strings.ToLower(*user.Email)

	if err := r.db.Model(model.User{}).Save(&model.User{
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

func (r *userRepository) Update(user *dto.CreateUserDto) error {
	email := strings.ToLower(*user.Email)
	user.Email = &email

	if err := r.db.Model(model.User{}).Save(model.User{
		ID:        0,
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

func (r *userRepository) FindByID(id int) (*dto.UserDto, error) {
	var user model.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
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

func (r *userRepository) FindAll(page, size int) ([]*dto.UserDto, error) {
	var users []*model.User

	if err := r.db.Offset((page - 1) * size).Limit(size).Find(&users).Error; err != nil {
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

func (r *userRepository) Delete(id int) error {
	if err := r.db.Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		log.Errorf("failed to delete user: %v", err)
		return err
	}

	return nil
}
