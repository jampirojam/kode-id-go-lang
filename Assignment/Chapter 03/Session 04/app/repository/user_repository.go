package repository

import (
	"ojam-test/c3/s4/app/entity"
)

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	FindByUserName(user entity.User) (entity.User, error)
	FindByUserId(userId int) (entity.User, error)
}