package request

import (
	"fmt"
)

func ErrorRequest(field, tag, param string) (message error) {
	switch tag {
	case "required":
		message = fmt.Errorf("%s is required", field)
	case "email":
		message = fmt.Errorf("%s must be valid email", field)
	case "lte":
		message = fmt.Errorf("%s must be equal or lower than %s", field, param)
	case "gte":
		message = fmt.Errorf("%s must be equal or greater than %s", field, param)
	default:
		message = fmt.Errorf("error to validate request")
	}

	return
}