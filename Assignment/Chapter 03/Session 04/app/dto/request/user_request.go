package request

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	UserName string `json:"user_name" validate:"required,gte=4,lte=25"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6,lte=20"`
	Age      int    `json:"age" validate:"required,gte=8,lte=80"`
}

type UserLoginRequest struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
