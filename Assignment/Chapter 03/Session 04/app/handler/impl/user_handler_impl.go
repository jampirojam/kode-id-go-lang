package impl

import (
	"net/http"
	"ojam-test/c3/s4/app/dto/request"
	"ojam-test/c3/s4/app/dto/response"
	"ojam-test/c3/s4/app/service"
	"ojam-test/c3/s4/app/util"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Svc service.UserService
}

func UserHandlerImpl(userService service.UserService) *UserHandler {
	return &UserHandler{
		Svc: userService,
	}
}

// @Summary      regist new user
// @Description  regist new user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request        body 	request.UserRequest		true	"user request is mandatory"
// @Success      200  			{array} response.UserResponse
// @Failure		 400			{array}	response.ErrorResponse	
// @Router       /v1/user/ 				[post]
func (uh *UserHandler) Register(ctx *gin.Context) {
	var req request.UserRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if err := util.ValidateJSON(ctx, req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Error: err,
		})
		return
	}

	res, err := uh.Svc.Register(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"context" : res,
	})
}

// @Summary      login
// @Description  regist new user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request        body 	request.UserLoginRequest	true	"user request is mandatory"
// @Success      200  			{array} response.UserLoginResponse
// @Failure		 400			{array}	response.ErrorResponse	
// @Router       /v1/user/login			[post]
func (uh *UserHandler) Login(ctx *gin.Context) {
	var req request.UserLoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if err := util.ValidateJSON(ctx, req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Error: err,
		})
		return
	}

	res, err := uh.Svc.Login(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"context" : res,
	})
}