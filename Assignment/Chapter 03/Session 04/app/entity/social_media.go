package entity

import (
	"time"
)

type SocialMedia struct {
	Id        int        `gorm:"primaryKey"`
	Name      string     `gorm:"not null; text"`
	Url       string     `gorm:"not null; text"`
	UserId    int        `gorm:"not null; integer"`
	CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time
	Deleted   bool `gorm:"default:false"`
}
