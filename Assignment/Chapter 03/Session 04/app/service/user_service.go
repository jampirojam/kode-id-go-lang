package service

import (
	"ojam-test/c3/s4/app/dto/request"
	"ojam-test/c3/s4/app/dto/response"
)

type UserService interface {
	Register(req request.UserRequest) (response.UserResponse, error)
	Login(req request.UserLoginRequest) (response.UserLoginResponse, error)
}