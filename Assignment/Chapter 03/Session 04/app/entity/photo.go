package entity

import (
	"time"
)

type Photo struct {
	Id        int        `gorm:"primaryKey"`
	Title     string     `gorm:"not null; text"`
	Caption   string     `gorm:"not null; text"`
	Url       string     `gorm:"not null; text"`
	UserId    int        `gorm:"not null; integer"`
	Comment	[]Comment
	CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time
	Deleted   bool `gorm:"default:false"`
}
