package entity

import "time"

type Book struct {
	Id             	int    			`gorm:"primaryKey"`
	Title          	string 			`gorm:"not null; text"`
	Author         	string
	Publisher      	string
	Year           	int
	BookCategoryId 	int 			`gorm:"not null"`
	CreatedAt      	*time.Time		`gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt		*time.Time		`gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt		*time.Time
	Deleted			bool			`gorm:"default:false"`
}

type BookCategory struct {
	Id				int				`gorm:"primaryKey"`
	Name  			string 			`gorm:"not null; unique; text"`
	Books []Book
	CreatedAt      	*time.Time		`gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt		*time.Time		`gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt		*time.Time
	Deleted			bool			`gorm:"default:false"`
}