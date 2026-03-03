package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/pkg/exception"
)

type customValidator struct {
	v *validator.Validate
}

func NewValidator() echo.Validator {
	return &customValidator{v: validator.New()}
}

func (cv *customValidator) Validate(i any) error {
	if err := cv.v.Struct(i); err != nil {
		return exception.BadRequest(err.Error())
	}
	return nil
}

func ValidateSchema(c *echo.Context, schema any) error {
	if err := c.Bind(schema); err != nil {
		return exception.BadRequest("invalid request: " + err.Error())
	}
	return c.Validate(schema)
}
