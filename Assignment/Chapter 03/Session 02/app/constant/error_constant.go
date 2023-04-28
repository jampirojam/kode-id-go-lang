package constant

// App Constant
const (
	PASSWORD_NOT_MATCH = "password not match"

	INVALID_ACCESS_TOKEN = "invalid access token"

	FAILED_TO_PARSE_TOKEN = "failed to parse token"

	FAILED_TO_PARSE_REFRESH_TOKEN = "failed to parse refresh token"

	USER_ROLE_AND_PRODUCT_TYPE_NOT_MATCH = "user role and product type not match"

	REQUEST_STRUCTURE_IS_NOT_VALID = "request structure is not valid"
)

// User Constant
const (
	USER_NOT_ALLOWED = "user not allowed to access"

	USER_EXIST = "user with username: %s already exist"

	USER_NOT_FOUND = "user with username: %s not found"

	USER_ROLE_NAME_NOT_FOUND = "user role name not found"
)

// Product Constant
const (
	CREATE_PRODUCT_ERROR = "there's an error when add new product"

	PRODUCT_NOT_AUTHORIZED = "product not authorized with user role"

	PRODUCT_TYPE_NAME_NOT_FOUND = "product type name not found"

	PRODUCT_NOT_MATCH = "product not match"

	PRODUCT_NOT_FOUND = "product with id: %d not found"

	PRODUCT_NOT_FOUND_OR_NOT_AUTHORIZED = "product with id: %d not found or not authorized with user role"

	PRODUCT_IS_EMPTY = "product is empty"
)
