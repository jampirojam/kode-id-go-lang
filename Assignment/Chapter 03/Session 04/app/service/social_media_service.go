package service

import (
	"ojam-test/c3/s4/app/dto/request"
	"ojam-test/c3/s4/app/dto/response"
)

type SocialMediaService interface {
	Add(req request.SocialMediaRequest, userId int) (response.SocialMediaResponse, error)
	UpdateById(id int, req request.SocialMediaRequest, userId int) (response.SocialMediaResponse, error)
	DeleteById(id, userId int) (response.SocialMediaResponse, error)
	GetById(id int, userId int) (response.SocialMediaResponse, error)
	GetAll(userId int) ([]response.SocialMediaResponse, error)
}