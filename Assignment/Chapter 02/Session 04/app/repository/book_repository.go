package repository

import (
	"ojam-test/c2/s4/app/entity"
)

type BookRepository interface {
	FindAll() ([]entity.Book, error)
	FindById(id int) (entity.Book, error)
	Create(book entity.Book) (entity.Book, error)
	UpdateById(id int, book entity.Book) (entity.Book, error)
	DeleteById(id int, book entity.Book) (entity.Book, error)
}