package response

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func Success(pctx *echo.Context, message string, result interface{}) error {
	return pctx.JSON(http.StatusOK, &Response{
		Status:  true,
		Message: message,
		Result:  result,
	})
}

func BadRequest(pctx *echo.Context, err error) error {
	return pctx.JSON(http.StatusBadRequest, &Response{
		Status:  false,
		Message: err.Error(),
		Result:  nil,
	})
}

func Unauthorized(pctx *echo.Context, err error) error {
	return pctx.JSON(http.StatusUnauthorized, &Response{
		Status:  false,
		Message: err.Error(),
		Result:  nil,
	})
}

func NotFound(pctx *echo.Context, err error) error {
	return pctx.JSON(http.StatusNotFound, &Response{
		Status:  false,
		Message: err.Error(),
		Result:  nil,
	})
}

func InternalError(pctx *echo.Context, err error) error {
	return pctx.JSON(http.StatusInternalServerError, &Response{
		Status:  false,
		Message: "Internal server error",
		Result:  err.Error(),
	})
}
