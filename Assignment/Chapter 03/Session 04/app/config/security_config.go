package config

import (
	"fmt"
	"net/http"
	"ojam-test/c3/s4/app/constant"
	"ojam-test/c3/s4/app/dto/response"
	"ojam-test/c3/s4/app/util"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(ctx *gin.Context) {
	_, err := ApiAuth(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	auth := ctx.GetHeader("Authorization")
	tokenAuth := strings.Split(auth, " ")[1]

	jwtToken, err := util.VerifyToken(tokenAuth)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	claims, success := jwtToken.Claims.(jwt.MapClaims)
	if !success {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Error: constant.FAILED_TO_PARSE_TOKEN,
		})
		return
	}

	ctx.Set("id", claims["id"])
	ctx.Next()
}

func UserAuth(ctx *gin.Context) int {

	userId, exist := ctx.Get("id")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusForbidden, response.ErrorResponse{
			Error: constant.USER_NOT_ALLOWED,
		})
		return 0
	}
	return int(userId.(float64))
}

func ApiAuth(ctx *gin.Context) (string, error) {
	auth := ctx.GetHeader("api-secret")
	if auth != constant.API_SECRET {
		return "", fmt.Errorf(constant.USER_NOT_ALLOWED)
	}

	return auth, nil
}