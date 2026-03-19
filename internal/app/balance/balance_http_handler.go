package balance

import (
	"errors"
	"log/slog"

	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/internal/response"
)

type BalanceHttpHandlerImpl struct {
	logger         *slog.Logger
	balanceUsecase BalanceUsecase
}

func NewBalanceHttpHandler(logger *slog.Logger, balanceUsecase BalanceUsecase) BalanceHttpHandler {
	return &BalanceHttpHandlerImpl{
		logger:         logger,
		balanceUsecase: balanceUsecase,
	}
}

func (h *BalanceHttpHandlerImpl) CoinTopUp(pctx *echo.Context) error {
	userId, err := h.getUserId(pctx)
	if err != nil {
		pctx.Logger().Error("failed to get user ID", "error", err)
		return response.BadRequest(pctx, err)
	}

	req := new(CoinTopUpRequest)
	if err := pctx.Bind(req); err != nil {
		pctx.Logger().Error("failed to bind request", "error", err)
		return response.BadRequest(pctx, err)
	}
	if err := pctx.Validate(req); err != nil {
		pctx.Logger().Error("failed to validate request", "error", err)
		return response.BadRequest(pctx, err)
	}

	req.UserId = userId
	newBalance, err := h.balanceUsecase.CoinAdding(req)
	if err != nil {
		pctx.Logger().Error("failed to add coin", "error", err)
		return response.InternalError(pctx, err)
	}

	return response.Success(pctx, "Coin added successfully", newBalance)
}

func (h *BalanceHttpHandlerImpl) CoinShowBalance(pctx *echo.Context) error {
	userId, err := h.getUserId(pctx)
	if err != nil {
		pctx.Logger().Error("failed to get user ID", "error", err)
		return response.BadRequest(pctx, err)
	}

	coin := h.balanceUsecase.CoinShowBalance(userId)

	return response.Success(pctx, "Coin shown successfully", coin)
}

func (h *BalanceHttpHandlerImpl) getUserId(pctx *echo.Context) (string, error) {
	userId, exists := pctx.Get("userId").(string)
	if !exists {
		return "", errors.New("user ID not found")
	}

	return userId, nil
}
