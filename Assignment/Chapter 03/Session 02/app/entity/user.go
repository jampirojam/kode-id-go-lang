package entity

import (
	"ojam-test/c3/s2/app/constant"
	"time"
)

type User struct {
	Id            int    `gorm:"primaryKey"`
	Name          string `gorm:"not null; text"`
	UserName      string `gorm:"not null; unique; text"`
	Password      string `gorm:"not null; text"`
	ProductTypeId int
	UserRoleId    int
	CreatedAt     *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt     *time.Time
	Deleted       bool `gorm:"default:false"`
}

type UserRole struct {
	Id        int                   `gorm:"primaryKey"`
	RoleType  constant.UserRoleType `gorm:"not null; unique; text"`
	Users     []User
	CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time
	Deleted   bool `gorm:"default:false"`
}
