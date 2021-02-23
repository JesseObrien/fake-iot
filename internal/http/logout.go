package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserLogOutHandler(tokenStore TokenStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		expireLoginToken(ctx, tokenStore)

		return ctx.JSON(http.StatusNoContent, "logged out")
	}
}
