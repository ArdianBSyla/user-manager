package model

import "time"

// User model represents a user in the system
type User struct {
	ID        int        `gorm:"primaryKey,autoIncrement"`
	CompanyID int        `gorm:"not null"`
	FirstName *string    `gorm:"not null"`
	LastName  *string    `gorm:"not null"`
	Email     *string    `gorm:"not null"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
}

func (u User) TableName() string {
	return "user"
}
