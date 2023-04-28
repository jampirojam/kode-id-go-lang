package entity

import (
	"time"
)

type Comment struct {
	Id        int        `gorm:"primaryKey"`
	Message   string     `gorm:"not null; text"`
	Url       string     `gorm:"not null; text"`
	UserId    int        `gorm:"not null; integer"`
	PhotoId   int        `gorm:"not null; integer"`
	CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time
	Deleted   bool `gorm:"default:false"`
}
