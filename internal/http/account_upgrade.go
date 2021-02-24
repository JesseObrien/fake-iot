package http

import (
	"net/http"

	"github.com/jesseobrien/fake-iot/internal/storage"
	"github.com/labstack/echo/v4"
)

type UpgradeAccountRequest struct {
	AccountId string `path:"id"`
}

func AccountUpgradeHandler(accountStore storage.AccountStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		upgradeReq := UpgradeAccountRequest{}

		if err := ctx.Bind(&upgradeReq); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid account id")
		}

		return ctx.JSON(http.StatusNoContent, "account upgraded")
	}
}
