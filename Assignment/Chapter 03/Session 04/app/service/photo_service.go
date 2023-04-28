package service

import (
	"ojam-test/c3/s4/app/dto/request"
	"ojam-test/c3/s4/app/dto/response"
)

type PhotoService interface {
	Add(req request.PhotoRequest, userId int) (response.PhotoResponse, error)
	UpdateById(id int, req request.PhotoRequest, userId int) (response.PhotoResponse, error)
	DeleteById(id, userId int) (response.PhotoResponse, error)
	GetById(id int) (response.PhotoResponse, error)
	GetAll() ([]response.PhotoResponse, error)
}