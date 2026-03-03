package validator

import (
	"net/http"
	"strings"

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
	if i == nil {
		return nil
	}
	if err := cv.v.Struct(i); err != nil {
		if valErr, ok := err.(validator.ValidationErrors); ok {
			messages := make([]string, 0, len(valErr))
			for _, e := range valErr {
				messages = append(messages, e.Field()+": "+e.Tag())
			}
			return echo.NewHTTPError(http.StatusBadRequest, strings.Join(messages, "; "))
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func ValidateSchema(c *echo.Context, schema any) error {
	if err := c.Bind(schema); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request: "+err.Error())
	}
	return c.Validate(schema)
}
