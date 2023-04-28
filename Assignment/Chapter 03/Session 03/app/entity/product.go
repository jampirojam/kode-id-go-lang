package entity

import (
	"ojam-test/c3/s2/app/constant"
	"time"
)

type Product struct {
	Id            int    `gorm:"primaryKey"`
	Name          string `gorm:"not null; text"`
	ProductTypeId int
	Price         int        `gorm:"not null; integer"`
	CreatedAt     *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt     *time.Time
	Deleted       bool `gorm:"default:false"`
}

type ProductType struct {
	Id        int                  `gorm:"primaryKey"`
	Type      constant.ProductType `gorm:"not null; unique; text"`
	Products  []Product
	Users     []User
	CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time
	Deleted   bool `gorm:"default:false"`
}
