package http

import (
	"fmt"
	"net/http"

	_ "github.com/jesseobrien/fake-iot/web/statik"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rakyll/statik/fs"
)

func Run(listenAddress, certPath, keyPath, apiToken string) error {
	e := echo.New()

	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	statikFS, err := fs.New()
	if err != nil {
		return fmt.Errorf("error initializing statik FS %w", err)
	}

	h := http.FileServer(statikFS)

	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", h)))

	e.POST("/metrics", IngestMetricsHandler(apiToken))

	return e.StartTLS(listenAddress, certPath, keyPath)
}
