package http

import (
	"crypto/subtle"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// IngestMetricsHandler will take user metrics in and store them into postgres
func IngestMetricsHandler(apiToken string) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Set the content-type to `application/json` instead of the default
		// `application/json;charset=utf-8` as the fakeiot cli doesn't like it
		ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		// Extract the bearer token
		authHeader := ctx.Request().Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Check that the apiToken is a valid match, otherwise error out
		if subtle.ConstantTimeCompare([]byte(token), []byte(apiToken)) == 0 {
			return ctx.JSON(http.StatusForbidden, errors.New("invalid authorization token"))
		}

		// @TODO make sure that the fakeiot `tests` pass on this handler (empty request, corrupted request, etc)
		// @TODO put together a custom type for the metrics payload
		// @TODO ensure the payload binds properly
		// @TODO write the metrics to postgres

		return ctx.JSON(http.StatusOK, "consumed metric")
	}
}
