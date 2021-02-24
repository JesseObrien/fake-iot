package http

import (
	"crypto/subtle"
	"errors"
	"net/http"

	"github.com/jesseobrien/fake-iot/internal/storage"
	"github.com/labstack/echo/v4"
)

// IngestMetricsHandler will take user metrics in and store them into postgres
func IngestMetricsHandler(apiToken string, accountStore storage.MetricStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Set the content-type to `application/json` instead of the default
		// `application/json;charset=utf-8` as the fakeiot cli doesn't like it
		ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		// Extract the bearer token
		token := getBearerTokenFromHeader(ctx)

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

		return ctx.JSON(http.StatusOK, "consumed metric")
	}
}
