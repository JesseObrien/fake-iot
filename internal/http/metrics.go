package http

import (
	"crypto/subtle"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/jesseobrien/fake-iot/internal/storage"
	"github.com/labstack/echo/v4"
)

// IngestMetricsHandler will take user metrics in and store them into postgres
func IngestMetricsHandler(apiToken string, accountStore storage.AccountStore) echo.HandlerFunc {
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

		metric := storage.UserLoginMetric{}

		if err := ctx.Bind(&metric); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "bad request")
		}

		if metric == (storage.UserLoginMetric{}) {
			return echo.NewHTTPError(http.StatusBadRequest, "metric should not be empty")
		}

		echoCtx := ctx.Request().Context()

		if err := accountStore.Write(echoCtx, metric); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// @TODO remove this, just testing it
		count, err := accountStore.CountByAccountId(echoCtx, metric.AccountID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		log.Printf("count of logins: %d", count)
		// @TODO remove the above

		return ctx.JSON(http.StatusOK, "consumed metric")
	}
}
