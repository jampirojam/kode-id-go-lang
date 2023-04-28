package impl

import (
	"errors"
	"fmt"
	"ojam-test/c3/s4/app/constant"
	"ojam-test/c3/s4/app/entity"
	"ojam-test/c3/s4/app/repository"
	"strconv"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func UserRepositoryImpl(db *gorm.DB) repository.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) Create(user entity.User) (entity.User, error) {
	newUser := entity.User{
		Name:          user.Name,
		UserName:      user.UserName,
		Password:      user.Password,
		Email: user.Email,
		Age: user.Age,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}

	err := ur.DB.Create(&newUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return entity.User{}, fmt.Errorf(constant.USER_EXIST, user.UserName)
		}

		return entity.User{}, err
	}

	return newUser, nil
}

func(ur *UserRepository) FindByUserName(user entity.User) (entity.User, error) {
	var findUser entity.User

	err := ur.DB.
		Where("user_name = ? AND deleted = ?", user.UserName, false).
		First(&findUser).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, fmt.Errorf(constant.USER_NOT_FOUND, user.UserName)
		}

		return entity.User{}, err
	}

	return findUser, nil
}

func(ur *UserRepository) FindByUserId(userId int) (entity.User, error) {
	var findUser entity.User

	err := ur.DB.
		Where("id = ? AND deleted = ?", userId, false).
		First(&findUser).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, fmt.Errorf(constant.USER_NOT_FOUND, strconv.Itoa(userId))
		}

		return entity.User{}, err
	}

	return findUser, nil
}
