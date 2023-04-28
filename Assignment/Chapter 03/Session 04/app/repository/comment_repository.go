package repository

import (
	"ojam-test/c3/s4/app/entity"
)

type CommentRepository interface {
	Create(comment entity.Comment, userId int) (entity.Comment, error)
	UpdateById(id int, comment entity.Comment, userId int) (entity.Comment, error)
	DeleteById(id int, comment entity.Comment, userId int) (entity.Comment, error)
	FindById(id int, userId int) (entity.Comment, error)
	FindAll(userId int) ([]entity.Comment, error)
	FindAllByPhotoId(photoId int) ([]entity.Comment, error)
}