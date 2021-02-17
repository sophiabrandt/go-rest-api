package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// validate holds the settings and caches for validating request struct values.
var validate *validator.Validate

func NewValidator() *validator.Validate {
	// Instantiate the validator for use.
	validate = validator.New()

	// Use JSON tag names for errors instead of Go struct names.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return validate
}

// decode unmarhals an incoming JSON request.
func decode(r *http.Request, val interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(val); err != nil {
		return err
	}
	return nil
}

// toErrResponse transforms the validation errors into a client-facing error response.
func toErrResponse(err error) *ErrorResponse {
	if fieldErrors, ok := err.(validator.ValidationErrors); ok {
		resp := ErrorResponse{
			Error:  "field validation error",
			Fields: make([]string, len(fieldErrors)),
		}

		for i, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				resp.Fields[i] = fmt.Sprintf("%s is a required field", err.Field())
			case "max":
				resp.Fields[i] = fmt.Sprintf("%s must be a maximum of %s in length", err.Field(), err.Param())
			case "url":
				resp.Fields[i] = fmt.Sprintf("%s must be a valid URL", err.Field())
			default:
				resp.Fields[i] = fmt.Sprintf("something wrong on %s; %s", err.Field(), err.Tag())
			}
		}

		return &resp
	}

	return nil
}
