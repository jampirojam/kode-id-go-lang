package impl

import (
	"errors"
	"fmt"
	"ojam-test/c3/s4/app/constant"
	"ojam-test/c3/s4/app/entity"
	"ojam-test/c3/s4/app/repository"

	"gorm.io/gorm"
)

type PhotoRepository struct {
	DB *gorm.DB
}

func PhotoRepositoryImpl(db *gorm.DB) repository.PhotoRepository {
	return &PhotoRepository{
		DB: db,
	}
}

func (pr *PhotoRepository) Create(photo entity.Photo, userId int) (entity.Photo, error) {
	newPhoto := entity.Photo{
		Title:     photo.Title,
		Caption:   photo.Caption,
		Url:       photo.Url,
		UserId:    userId,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
	}

	err := pr.DB.Create(&newPhoto).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return entity.Photo{}, fmt.Errorf(constant.CREATE_PHOTO_ERROR)
		}

		return entity.Photo{}, err
	}

	return newPhoto, nil
}

func (pr *PhotoRepository) UpdateById(id int, photo entity.Photo, userId int) (entity.Photo, error) {
	updatedPhoto := entity.Photo{}

	err := pr.DB.
		Where("id = ? AND user_id = ? AND deleted = ?", id, userId, false).
		First(&updatedPhoto).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Photo{}, fmt.Errorf(constant.PHOTO_NOT_FOUND, id)
		}

		return entity.Photo{}, err
	}

	err = pr.DB.Model(&updatedPhoto).
		Where("id = ? AND user_id = ? AND deleted = ?", id, userId, false).
		Updates(entity.Photo{
			Title:     photo.Title,
			Caption:   photo.Caption,
			Url:       photo.Url,
			UserId:    userId,
			UpdatedAt: photo.UpdatedAt,
		}).Error

	if err != nil {
		return entity.Photo{}, err
	}

	return updatedPhoto, nil
}

func (pr *PhotoRepository) DeleteById(id int, photo entity.Photo, userId int) (entity.Photo, error) {
	deletedPhoto := entity.Photo{}

	err := pr.DB.
		Where("id = ? AND user_id = ? AND deleted = ?", id, userId, false).
		First(&deletedPhoto).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Photo{}, fmt.Errorf(constant.PHOTO_NOT_FOUND, id)
		}

		return entity.Photo{}, err
	}

	err = pr.DB.Model(&deletedPhoto).
		Where("id = ? AND user_id = ? AND deleted = ?", id, userId, false).
		Updates(entity.Photo{
			UpdatedAt: photo.UpdatedAt,
			DeletedAt: photo.DeletedAt,
			Deleted:   photo.Deleted,
		}).Error

	if err != nil {
		return entity.Photo{}, err
	}

	return deletedPhoto, nil
}

func (pr *PhotoRepository) FindById(id int) (entity.Photo, error) {
	findPhoto := entity.Photo{}

	err := pr.DB.
		Where("id = ? AND deleted = ?", id, false).
		First(&findPhoto).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Photo{}, fmt.Errorf(constant.PHOTO_NOT_FOUND, id)
		}

		return entity.Photo{}, err
	}

	return findPhoto, nil
}


func (pr *PhotoRepository) FindAll() ([]entity.Photo, error) {
	photoList := []entity.Photo{}

	err := pr.DB.
		Where("deleted = ?", false).
		Find(&photoList).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.Photo{}, fmt.Errorf(constant.PHOTO_IS_EMPTY)
		}

		return []entity.Photo{}, err
	}

	return photoList, nil
}
