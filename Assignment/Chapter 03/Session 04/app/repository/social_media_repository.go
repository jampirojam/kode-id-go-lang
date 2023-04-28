package repository

import (
	"ojam-test/c3/s4/app/entity"
)

type SocialMediaRepository interface {
	Create(socialMedia entity.SocialMedia, userId int) (entity.SocialMedia, error)
	UpdateById(id int, socialMedia entity.SocialMedia, userId int) (entity.SocialMedia, error)
	DeleteById(id int, socialMedia entity.SocialMedia, userId int) (entity.SocialMedia, error)
	FindById(id int, userId int) (entity.SocialMedia, error)
	FindAll(userId int) ([]entity.SocialMedia, error)
}