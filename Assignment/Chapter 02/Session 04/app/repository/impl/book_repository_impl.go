package impl

import (
	"errors"
	"fmt"
	"ojam-test/c2/s4/app/entity"
	"ojam-test/c2/s4/app/repository"

	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) repository.BookRepository {
	return &BookRepository{
		DB: db,
	}
}

func (repository *BookRepository) FindAll() ([]entity.Book, error) {
	var books []entity.Book

	err := repository.DB.
		Where("deleted = ?", false).
		Find(&books).
		Error

	if err != nil {
		return []entity.Book{}, err
	}

	return books, nil
}

func (repository *BookRepository) FindById(id int) (entity.Book, error) {
	var book entity.Book

	err := repository.DB.
		Where("id = ? AND deleted = ?", id, false).
		First(&book, "id = ?", id).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Book{}, fmt.Errorf("book with Id %d not found", id)
		}

		return entity.Book{}, err
	}

	return book, nil
}

func (repository *BookRepository) Create(book entity.Book) (entity.Book, error) {
	newBook := entity.Book{
		Title:     book.Title,
		Author:    book.Author,
		Publisher: book.Publisher,
		Year:      book.Year,
		BookCategoryId: book.BookCategoryId,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}

	err := repository.DB.Create(&newBook).Error
	if err != nil {
		return entity.Book{}, err
	}

	return newBook, nil
}

func (repository *BookRepository) UpdateById(id int, book entity.Book) (entity.Book, error) {
	updatedBook := entity.Book{}

	err := repository.DB.
		Where("id = ? AND deleted = ?", id, false).
		First(&updatedBook).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Book{}, fmt.Errorf("book with Id: %d not found", id)
		}

		return entity.Book{}, err
	}

	err = repository.DB.Model(&updatedBook).Updates(
		entity.Book{
			Title:          book.Title,
			Author:         book.Author,
			Publisher:      book.Publisher,
			Year:           book.Year,
			BookCategoryId: book.BookCategoryId,
		},
	).Error

	if err != nil {
		return entity.Book{}, err
	}

	return updatedBook, nil
}

func (repository *BookRepository) DeleteById(id int, book entity.Book) (entity.Book, error) {
	deletedBook := entity.Book{}

	err := repository.DB.
		Where("id = ? AND deleted = ?", id, false).
		First(&deletedBook).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Book{}, fmt.Errorf("data with Id: %d not found", id)
		}

		return entity.Book{}, err
	}

	err = repository.DB.Model(&deletedBook).Updates(
		entity.Book{
			UpdatedAt: book.UpdatedAt,
			DeletedAt: book.DeletedAt,
			Deleted:   book.Deleted,
		},
	).Error

	if err != nil {
		return entity.Book{}, err
	}

	return deletedBook, nil
}
