package request_validator

import (
	"strings"

	"github.com/herryg91/kless_go_lib/http/response"

	"github.com/mcuadros/go-defaults"
	"gopkg.in/validator.v2"
)

func Validate(req interface{}) *response.Response {
	defaults.SetDefaults(req)

	other_errors := []response.ErrorFieldResponse{}
	if errs := validator.Validate(req); errs != nil {
		for field, err := range errs.(validator.ErrorMap) {
			other_errors = append(other_errors, response.ErrorFieldResponse{
				Field:   field,
				Message: strings.ReplaceAll(err.Error(), "{field}", field),
			})
		}
	}

	if len(other_errors) > 0 {
		return &response.Response{
			Success:     false,
			Message:     "",
			ErrorCode:   "REQUEST_VALIDATION",
			ProcessTime: "0ms",
			Data:        other_errors,
		}
	}
	return nil
}
