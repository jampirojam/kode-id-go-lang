package service

import (
	"ojam-test/c3/s2/app/dto/request"
	"ojam-test/c3/s2/app/dto/response"
)

type UserService interface {
	Register(req request.UserRequest) (response.UserResponse, error)
	Login(req request.UserLoginRequest) (response.UserLoginResponse, error)
	Relog(userId int) (response.UserLoginResponse, error)
}