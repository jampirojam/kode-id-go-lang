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

// Photo Constant
const (
	CREATE_PHOTO_ERROR = "there's an error when add new photo"

	PHOTO_NOT_FOUND = "photo with id: %d not found"

	PHOTO_NOT_FOUND_OR_NOT_AUTHORIZED = "photo with id: %d not found in user with id: %d"

	PHOTO_IS_EMPTY = "photo is empty"
)

// Photo Constant
const (
	CREATE_COMMENT_ERROR = "there's an error when add new comment"

	COMMENT_NOT_FOUND = "comment with id: %d not found"

	COMMENT_NOT_FOUND_OR_NOT_AUTHORIZED = "comment with id: %d not found in user with id: %d"

	COMMENT_IS_EMPTY = "comment is empty"
)

// Social Media Constant
const (
	CREATE_SOCIAL_MEDIA_ERROR = "there's an error when add new comment"

	SOCIAL_MEDIA_NOT_FOUND = "comment with id: %d not found"

	SOCIAL_MEDIA_NOT_FOUND_OR_NOT_AUTHORIZED = "comment with id: %d not found in user with id: %d"

	SOCIAL_MEDIA_IS_EMPTY = "comment is empty"
)
