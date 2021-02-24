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

		if err := echo.PathParamsBinder(ctx).MustString("id", &upgradeReq.AccountId).BindError(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid account id")
		}

		if err := accountStore.Upgrade(ctx.Request().Context(), upgradeReq.AccountId, storage.AccountPlanTypeEnterprise); err != nil {
			ctx.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "error upgrading account")
		}

		newAccountInfo, err := accountStore.GetAccountInfo(ctx.Request().Context(), upgradeReq.AccountId)

		if err != nil {
			ctx.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "error upgrading account")
		}

		return ctx.JSON(http.StatusOK, newAccountInfo)
	}
}
