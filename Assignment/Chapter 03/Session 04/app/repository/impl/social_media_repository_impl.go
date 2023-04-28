package impl

import (
	"errors"
	"fmt"
	"ojam-test/c3/s4/app/constant"
	"ojam-test/c3/s4/app/entity"
	"ojam-test/c3/s4/app/repository"

	"gorm.io/gorm"
)

type SocialMediaRepository struct {
	DB *gorm.DB
}

func SocialMediaRepositoryImpl(db *gorm.DB) repository.SocialMediaRepository {
	return &SocialMediaRepository{
		DB: db,
	}
}

func (pr *SocialMediaRepository) Create(socialMedia entity.SocialMedia, userId int) (entity.SocialMedia, error) {
	newSocialMedia := entity.SocialMedia{
		Name:      socialMedia.Name,
		Url:       socialMedia.Url,
		UserId:    userId,
		CreatedAt: socialMedia.CreatedAt,
		UpdatedAt: socialMedia.UpdatedAt,
	}

	err := pr.DB.Create(&newSocialMedia).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return entity.SocialMedia{}, fmt.Errorf(constant.CREATE_COMMENT_ERROR)
		}

		return entity.SocialMedia{}, err
	}

	return newSocialMedia, nil
}

func (pr *SocialMediaRepository) UpdateById(id int, socialMedia entity.SocialMedia, userId int) (entity.SocialMedia, error) {
	updatedSocialMedia := entity.SocialMedia{}

	err := pr.DB.
		Where("id = ? AND user_id = ? AND deleted = ?", id, userId, false).
		First(&updatedSocialMedia).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.SocialMedia{}, fmt.Errorf(constant.COMMENT_NOT_FOUND, id)
		}

		return entity.SocialMedia{}, err
	}

	err = pr.DB.Model(&updatedSocialMedia).
		Where("id = ? AND user_id = ? AND deleted = ?", id, userId, false).
		Updates(entity.SocialMedia{
			Name:      socialMedia.Name,
			Url:       socialMedia.Url,
			UserId:    userId,
			UpdatedAt: socialMedia.UpdatedAt,
		}).Error

	if err != nil {
		return entity.SocialMedia{}, err
	}

	return updatedSocialMedia, nil
}

func (pr *SocialMediaRepository) DeleteById(id int, socialMedia entity.SocialMedia, userId int) (entity.SocialMedia, error) {
	deletedSocialMedia := entity.SocialMedia{}

	err := pr.DB.
		Where("id = ? AND user_id = ? AND deleted = ?", id, userId, false).
		First(&deletedSocialMedia).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.SocialMedia{}, fmt.Errorf(constant.COMMENT_NOT_FOUND, id)
		}

		return entity.SocialMedia{}, err
	}

	err = pr.DB.Model(&deletedSocialMedia).
		Where("id = ? AND user_id = ? AND deleted = ?", id, userId, false).
		Updates(entity.SocialMedia{
			UpdatedAt: socialMedia.UpdatedAt,
			DeletedAt: socialMedia.DeletedAt,
			Deleted:   socialMedia.Deleted,
		}).Error

	if err != nil {
		return entity.SocialMedia{}, err
	}

	return deletedSocialMedia, nil
}

func (pr *SocialMediaRepository) FindById(id, userId int) (entity.SocialMedia, error) {
	findSocialMedia := entity.SocialMedia{}

	err := pr.DB.
		Where("id = ? AND user_id = ? AND deleted = ?", id, userId, false).
		First(&findSocialMedia).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.SocialMedia{}, fmt.Errorf(constant.COMMENT_NOT_FOUND, id)
		}

		return entity.SocialMedia{}, err
	}

	return findSocialMedia, nil
}

func (pr *SocialMediaRepository) FindAll(userId int) ([]entity.SocialMedia, error) {
	socialMediaList := []entity.SocialMedia{}

	err := pr.DB.
		Where("user_id = ? AND deleted = ?", userId, false).
		Find(&socialMediaList).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.SocialMedia{}, fmt.Errorf(constant.COMMENT_IS_EMPTY)
		}

		return []entity.SocialMedia{}, err
	}

	return socialMediaList, nil
}

