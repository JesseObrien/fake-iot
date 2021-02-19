package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserLogOutHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := expireLoginToken(ctx)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "error logging out")
		}

		ctx.SetCookie(cookie)

		return ctx.JSON(http.StatusNoContent, "logged out")
	}
}
