package service

import (
	"ojam-test/c3/s4/app/dto/request"
	"ojam-test/c3/s4/app/dto/response"
)

type CommentService interface {
	Add(req request.CommentRequest, userId int) (response.CommentResponse, error)
	UpdateById(id int, req request.CommentRequest, userId int) (response.CommentResponse, error)
	DeleteById(id, userId int) (response.CommentResponse, error)
	GetById(id int, userId int) (response.CommentResponse, error)
	GetAll(userId int) ([]response.CommentResponse, error)
}