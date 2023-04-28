package repository

import (
	"ojam-test/c3/s4/app/entity"
)

type PhotoRepository interface {
	Create(photo entity.Photo, userId int) (entity.Photo, error)
	UpdateById(id int, photo entity.Photo, userId int) (entity.Photo, error)
	DeleteById(id int, photo entity.Photo, userId int) (entity.Photo, error)
	FindById(id int) (entity.Photo, error)
	FindAll() ([]entity.Photo, error)
}