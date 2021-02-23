package http

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/jesseobrien/fake-iot/internal/storage"
	_ "github.com/jesseobrien/fake-iot/web/statik"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rakyll/statik/fs"
)

func Run(database *sql.DB, listenAddress, certPath, keyPath, apiToken string) error {
	e := echo.New()

	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	statikFS, err := fs.New()
	if err != nil {
		return fmt.Errorf("error initializing statik FS %w", err)
	}

	h := http.FileServer(statikFS)

	// New up the account store with a database connection
	accountStore := storage.NewPgAccountStore(database)

	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", h)))

	e.POST("/metrics", IngestMetricsHandler(apiToken, accountStore))

	return e.StartTLS(listenAddress, certPath, keyPath)
}
