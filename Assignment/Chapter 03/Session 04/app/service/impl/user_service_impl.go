package impl

import (
	"fmt"
	"ojam-test/c3/s4/app/constant"
	"ojam-test/c3/s4/app/dto/request"
	"ojam-test/c3/s4/app/dto/response"
	"ojam-test/c3/s4/app/entity"
	"ojam-test/c3/s4/app/repository"
	"ojam-test/c3/s4/app/util"
	"time"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func UserServiceImpl(userRepository repository.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepository,
	}
}

func (svc *UserService) Register(req request.UserRequest) (response.UserResponse, error) {
	setTime := time.Now()
	password, err := util.EncryptPassword(req.Password)
	if err != nil {
		return response.UserResponse{}, err
	}

	newUser := entity.User{
		Name:      req.Name,
		Email:     req.Email,
		UserName:  req.UserName,
		Password:  password,
		Age:       req.Age,
		CreatedAt: &setTime,
		UpdatedAt: &setTime,
	}

	res, err := svc.UserRepo.Create(newUser)
	if err != nil {
		return response.UserResponse{}, err
	}

	return generateUserResponse(&res)
}

func (svc *UserService) Login(req request.UserLoginRequest) (response.UserLoginResponse, error) {
	userLogin := entity.User{
		UserName: req.UserName,
		Password: req.Password,
	}

	res, err := svc.UserRepo.FindByUserName(userLogin)
	if err != nil {
		return response.UserLoginResponse{}, err
	}

	if !util.CheckPassword(res.Password, req.Password) {
		return response.UserLoginResponse{}, fmt.Errorf(constant.PASSWORD_NOT_MATCH)
	}

	accessToken, err := util.GenerateToken(res)
	if err != nil {
		return response.UserLoginResponse{}, err
	}

	return response.UserLoginResponse{
		Id:          res.Id,
		Name:        res.Name,
		AccessToken: accessToken,
	}, nil
}

func generateUserResponse(user *entity.User) (response.UserResponse, error) {
	return response.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		UserName:  user.UserName,
		Email:     user.Email,
		Age:       user.Age,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		Deleted:   user.Deleted,
	}, nil
}
