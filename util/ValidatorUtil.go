package util

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"majooTest/helpers"
	"net/http"
)

func ValidationBindJsonField(err error) interface{} {
	var ve validator.ValidationErrors
	var response helpers.Response
	if errors.As(err, &ve) {
		out := make([]helpers.AppError, len(ve))
		for i, fe := range ve {
			out[i] = helpers.AppError{http.StatusBadRequest, getErrorMsg(fe)}
			response = helpers.ResponseError(getErrorMsg(fe), http.StatusBadRequest)
		}
	} else {
		response = helpers.ResponseError(err.Error(), http.StatusBadRequest)
	}
	return response
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required " + fe.Field()
	case "lte":
		return "Should be less than" + fe.Field()
	case "gte":
		return "Should be greater than" + fe.Field()
	case "email":
		return "failed email format"
	case "unique":
		return fe.Field() + "already exist"
	}
	return "Unknown error"
}
