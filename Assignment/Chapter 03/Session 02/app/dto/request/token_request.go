package request

import "github.com/golang-jwt/jwt/v4"

type TokenRequest struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	UserName      string `json:"user_name"`
	ProductTypeId int    `json:"product_type_id"`
	UserRoleId    int    `json:"user_role_id"`
	jwt.StandardClaims
}
