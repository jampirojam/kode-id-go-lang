package request

import "github.com/golang-jwt/jwt/v4"

type TokenRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	jwt.StandardClaims
}
