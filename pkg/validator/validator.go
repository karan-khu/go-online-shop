package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

type customValidator struct {
	v *validator.Validate
}

func NewValidator() echo.Validator {
	return &customValidator{v: validator.New()}
}

func (cv *customValidator) Validate(i any) error {
	if err := cv.v.Struct(i); err != nil {
		return err
	}
	return nil
}

func ValidateSchema(c *echo.Context, schema any) error {
	if err := c.Bind(schema); err != nil {
		return err
	}
	return c.Validate(schema)
}
