package model

import "time"

// Company represents a company in the system
type Company struct {
	ID        int        `gorm:"primaryKey,autoIncrement"`
	Name      *string    `gorm:"not null"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
}

func (c Company) TableName() string {
	return "company"
}
