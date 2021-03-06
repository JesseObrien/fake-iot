package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jesseobrien/fake-iot/internal/storage"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIngestMetrics(t *testing.T) {
	testAccountId := "47f3c307-6344-49e7-961c-ea200e950a89"
	testUserId := "cbb3710e-8fa4-4f4e-9114-f27170887b16"
	testPayload := fmt.Sprintf(`{"account_id":"%s","user_id":"%s","timestamp":"2021-02-19T15:33:10.737127483Z"}`, testAccountId, testUserId)

	metricStore := storage.NewMemMetricStore()
	accountUpdateStore := storage.NewAccountUpdateStore()

	testApiToken := "882e8f9b-76a3-46fb-9f7e-bd536bdf5795"

	handler := IngestMetricsHandler(testApiToken, metricStore, accountUpdateStore)

	// Set up the request context
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/metrics", strings.NewReader(testPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", testApiToken))

	rec := httptest.NewRecorder()
	echoCtx := e.NewContext(req, rec)

	require.NoError(t, handler(echoCtx))

	assert.Equal(t, http.StatusOK, rec.Code)
	written, err := metricStore.WroteMetric(testAccountId, testUserId)
	assert.NoError(t, err)
	assert.True(t, written)
}
