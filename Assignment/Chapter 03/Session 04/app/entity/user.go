package entity

import (
	"time"
)

type User struct {
	Id          int    `gorm:"primaryKey"`
	Name        string `gorm:"not null; text"`
	UserName    string `gorm:"not null; unique; text"`
	Email       string `gorm:"not null; unique; text"`
	Password    string `gorm:"not null; text"`
	Age         int    `gorm:"not null; integer"`
	Photos      []Photo
	Comments    []Comment
	SocialMedia []SocialMedia
	CreatedAt   *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time
	Deleted     bool `gorm:"default:false"`
}
