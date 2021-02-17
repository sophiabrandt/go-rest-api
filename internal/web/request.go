package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/go-playground/validator/v10"
)

const (
	alphaSpaceRegexString string = "^[a-zA-Z .]+$"
	dateRegexString       string = "^(((19|20)([2468][048]|[13579][26]|0[48])|2000)[/-]02[/-]29|((19|20)[0-9]{2}[/-](0[469]|11)[/-](0[1-9]|[12][0-9]|30)|(19|20)[0-9]{2}[/-](0[13578]|1[02])[/-](0[1-9]|[12][0-9]|3[01])|(19|20)[0-9]{2}[/-]02[/-](0[1-9]|1[0-9]|2[0-8])))$"
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

	validate.RegisterValidation("alpha_space", isAlphaSpace)
	validate.RegisterValidation("date", isDate)

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

func isAlphaSpace(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(alphaSpaceRegexString)
	return reg.MatchString(fl.Field().String())
}

func isDate(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(dateRegexString)
	return reg.MatchString(fl.Field().String())
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
			case "alpha_space":
				resp.Fields[i] = fmt.Sprintf("%s can only contain alphabetic and space characters", err.Field())
			case "date":
				resp.Fields[i] = fmt.Sprintf("%s must be a valid date in the format: 2006-01-02", err.Field())
			default:
				resp.Fields[i] = fmt.Sprintf("something wrong on %s; %s", err.Field(), err.Tag())
			}
		}

		return &resp
	}

	return nil
}

// params returns the web call parameters from the request.
func params(r *http.Request) map[string]string {
	return httptreemux.ContextParams(r.Context())
}
