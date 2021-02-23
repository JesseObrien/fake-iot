package http

import (
	"net/http"

	"github.com/jesseobrien/fake-iot/internal/storage"
	"github.com/labstack/echo/v4"
)

func UserLogOutHandler(tokenStore *storage.TokenStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := getBearerTokenFromHeader(ctx)
		tokenStore.Expire(token)

		return ctx.JSON(http.StatusNoContent, "logged out")
	}
}
