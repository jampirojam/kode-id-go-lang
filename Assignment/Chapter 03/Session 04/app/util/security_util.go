package util

import (
	"fmt"
	"ojam-test/c3/s4/app/constant"
	"ojam-test/c3/s4/app/dto/request"
	"ojam-test/c3/s4/app/entity"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func CheckPassword(actualPassword, requestPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(actualPassword), []byte(requestPassword))

	return err == nil
}

func GenerateToken(user entity.User) (string, error) {
	tokenRequest := request.TokenRequest{
		Id:       user.Id,
		Name:     user.Name,
		UserName: user.UserName,
		Email:    user.Email,
		Age:      user.Age,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: constant.TOKEN_LIFESPAN,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenRequest)

	token, err := jwtToken.SignedString([]byte(constant.KEY_SECRET))
	return token, err
}

func VerifyToken(token string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, success := t.Method.(*jwt.SigningMethodHMAC)
		if !success {
			return nil, fmt.Errorf(constant.INVALID_ACCESS_TOKEN)
		}

		return []byte(constant.KEY_SECRET), nil
	})

	return jwtToken, err
}
