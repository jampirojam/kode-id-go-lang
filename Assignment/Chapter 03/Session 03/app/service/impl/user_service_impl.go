package impl

import (
	"fmt"
	"ojam-test/c3/s2/app/constant"
	"ojam-test/c3/s2/app/dto/request"
	"ojam-test/c3/s2/app/dto/response"
	"ojam-test/c3/s2/app/entity"
	"ojam-test/c3/s2/app/repository"
	"ojam-test/c3/s2/app/util"
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

	if req.UserRoleName != string(constant.UserRole.HEAD) {
		if !util.CheckSuitabilityProductTypeAndUserRole(req.ProductTypeName, req.UserRoleName) {
			return response.UserResponse{}, fmt.Errorf(constant.USER_ROLE_AND_PRODUCT_TYPE_NOT_MATCH)
		}
	}

	typeId, err := util.GetProductTypeId(req.ProductTypeName)
	if err != nil {
		return response.UserResponse{}, err
	}

	roleId, err := util.GetUserRoleId(req.UserRoleName)
	if err != nil {
		return response.UserResponse{}, err
	}

	newUser := entity.User{
		Name:          req.Name,
		UserName:      req.UserName,
		Password:      password,
		ProductTypeId: typeId,
		UserRoleId:    roleId,
		CreatedAt:     &setTime,
		UpdatedAt:     &setTime,
	}

	res, err := svc.UserRepo.Create(newUser)
	if err != nil {
		return response.UserResponse{}, err
	}

	return generateUserResponse(&res)
}

func (svc *UserService) Login(req request.UserLoginRequest) (response.UserLoginResponse, error) {
	userLogin := entity.User{
		UserName:      req.UserName,
		Password:      req.Password,
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

	refreshToken, err := util.GenerateRefreshToken(res)
	if err != nil {
		return response.UserLoginResponse{}, err
	}

	return response.UserLoginResponse{
		Id: res.Id,
		Name: res.Name,
		ProductTypeId: res.ProductTypeId,
		UserRoleId: res.UserRoleId,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (svc *UserService) Relog(userId int) (response.UserLoginResponse, error) {
	res, err := svc.UserRepo.FindByUserId(userId)
	if err != nil {
		return response.UserLoginResponse{}, err
	}

	accessToken, err := util.GenerateToken(res)
	if err != nil {
		return response.UserLoginResponse{}, err
	}

	refreshToken, err := util.GenerateRefreshToken(res)
	if err != nil {
		return response.UserLoginResponse{}, err
	}

	return response.UserLoginResponse{
		Id: res.Id,
		Name: res.Name,
		ProductTypeId: res.ProductTypeId,
		UserRoleId: res.UserRoleId,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func generateUserResponse(user *entity.User) (response.UserResponse, error) {
	return response.UserResponse{
		Id:            user.Id,
		Name:          user.Name,
		UserName:      user.UserName,
		Password:      user.Password,
		ProductTypeId: user.ProductTypeId,
		UserRoleId:    user.UserRoleId,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
		DeletedAt:     user.DeletedAt,
		Deleted:       user.Deleted,
	}, nil
}
