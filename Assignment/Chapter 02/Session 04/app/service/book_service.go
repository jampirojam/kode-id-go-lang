package service

import (
	"ojam-test/c2/s4/app/dto/request"
	"ojam-test/c2/s4/app/dto/response"
)

type BookService interface {
	GetAll() ([]response.BookResponse, error)
	UpdateById(id int, req request.BookRequest) (response.BookResponse, error)
	GetById(id int) (response.BookResponse, error)
	DeleteById(id int) (response.BookResponse, error)
	Create(req request.BookRequest) (response.BookResponse, error)
}