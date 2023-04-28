package impl

import (
	"errors"
	"fmt"
	"ojam-test/c3/s4/app/constant"
	"ojam-test/c3/s4/app/entity"
	"ojam-test/c3/s4/app/repository"

	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func CommentRepositoryImpl(db *gorm.DB) repository.CommentRepository {
	return &CommentRepository{
		DB: db,
	}
}

func (pr *CommentRepository) Create(comment entity.Comment, userId int) (entity.Comment, error) {
	newComment := entity.Comment{
		Message:   comment.Message,
		Url:       comment.Url,
		PhotoId:   comment.PhotoId,
		UserId:    userId,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}

	err := pr.DB.Create(&newComment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return entity.Comment{}, fmt.Errorf(constant.CREATE_PHOTO_ERROR)
		}

		return entity.Comment{}, err
	}

	return newComment, nil
}

func (pr *CommentRepository) UpdateById(id int, comment entity.Comment, userId int) (entity.Comment, error) {
	updatedComment := entity.Comment{}

	err := pr.DB.
		Where("id = ? AND user_id = ? AND deleted = ?", id, userId, false).
		First(&updatedComment).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Comment{}, fmt.Errorf(constant.PHOTO_NOT_FOUND, id)
		}

		return entity.Comment{}, err
	}

	err = pr.DB.Model(&updatedComment).
		Where("id = ? AND user_id = ? AND deleted = ?", id, userId, false).
		Updates(entity.Comment{
			Message:   comment.Message,
			Url:       comment.Url,
			PhotoId:   comment.PhotoId,
			UserId:    userId,
			UpdatedAt: comment.UpdatedAt,
		}).Error

	if err != nil {
		return entity.Comment{}, err
	}

	return updatedComment, nil
}

func (pr *CommentRepository) DeleteById(id int, comment entity.Comment, userId int) (entity.Comment, error) {
	deletedComment := entity.Comment{}

	err := pr.DB.
		Where("id = ? AND user_id = ? AND deleted = ?", id, userId, false).
		First(&deletedComment).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Comment{}, fmt.Errorf(constant.PHOTO_NOT_FOUND, id)
		}

		return entity.Comment{}, err
	}

	err = pr.DB.Model(&deletedComment).
		Where("id = ? AND user_id = ? AND deleted = ?", id, userId, false).
		Updates(entity.Comment{
			UpdatedAt: comment.UpdatedAt,
			DeletedAt: comment.DeletedAt,
			Deleted:   comment.Deleted,
		}).Error

	if err != nil {
		return entity.Comment{}, err
	}

	return deletedComment, nil
}

func (pr *CommentRepository) FindById(id, userId int) (entity.Comment, error) {
	findComment := entity.Comment{}

	err := pr.DB.
		Where("id = ? AND user_id = ? AND deleted = ?", id, userId, false).
		First(&findComment).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Comment{}, fmt.Errorf(constant.PHOTO_NOT_FOUND, id)
		}

		return entity.Comment{}, err
	}

	return findComment, nil
}

func (pr *CommentRepository) FindAll(userId int) ([]entity.Comment, error) {
	commentList := []entity.Comment{}

	err := pr.DB.
		Where("user_id = ? AND deleted = ?", userId, false).
		Find(&commentList).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.Comment{}, fmt.Errorf(constant.PHOTO_IS_EMPTY)
		}

		return []entity.Comment{}, err
	}

	return commentList, nil
}

func (pr *CommentRepository) FindAllByPhotoId(photoId int) ([]entity.Comment, error) {
	commentList := []entity.Comment{}

	err := pr.DB.
		Where("photo_id = ? AND deleted = ?", photoId, false).
		Find(&commentList).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.Comment{}, fmt.Errorf(constant.PHOTO_IS_EMPTY)
		}

		return []entity.Comment{}, err
	}

	return commentList, nil
}