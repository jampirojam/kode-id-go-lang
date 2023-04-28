package response

import "time"

type UserResponse struct {
	Id            int        `json:"id"`
	Name          string     `json:"name"`
	UserName      string     `json:"user_name"`
	Email         string     `json:"email"`
	Password      string     `json:"password"`
	Age           int        `json:"age"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	Deleted       bool       `json:"deleted"`
}

type UserLoginResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	AccessToken  string `json:"access_token"`
}
