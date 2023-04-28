package request

type UserRequest struct {
	Name            string `json:"name"`
	UserName        string `json:"user_name"`
	Password        string `json:"password"`
	ProductTypeName string `json:"product_type"`
	UserRoleName    string `json:"user_role"`
}

type UserLoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
