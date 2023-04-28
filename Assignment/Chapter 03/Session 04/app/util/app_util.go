package util

import (
	"ojam-test/c3/s4/app/dto/request"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

func Validator(req interface{}) []error {
	validate := validator.New()

	err := validate.Struct(req)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, request.ErrorRequest(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}

func ValidateJSON(ctx *gin.Context, req interface{}) []string {
	validateErr := Validator(req)
	if validateErr != nil {
		errResponse := make([]string, len(validateErr))
		for i, err := range validateErr {
			errResponse[i] = err.Error()
		}

		return errResponse
	}

	return nil
}

func GenerateUUID() string {
	id := uuid.New()

	return id.String()
}