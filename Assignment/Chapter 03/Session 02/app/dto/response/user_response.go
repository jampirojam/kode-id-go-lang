package response

import "time"

type UserResponse struct {
	Id            int        `json:"id"`
	Name          string     `json:"name"`
	UserName      string     `json:"user_name"`
	Password      string     `json:"password"`
	ProductTypeId int        `json:"product_type_id"`
	UserRoleId    int        `json:"user_role_id"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	Deleted       bool       `json:"deleted"`
}

type UserLoginResponse struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	ProductTypeId int    `json:"product_type_id"`
	UserRoleId    int    `json:"user_role_id"`
	AccessToken   string `json:"access_token"`
	RefreshToken  string `json:"refresh_token"`
}
