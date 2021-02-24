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

// As talked about in the design doc, I'm hard coding this. Ideally it would be stored
// in the database and looked up on request
const UserEmail = "test@example.com"

// As in the design doc, I'm hard coding the password and will hash it on startup. Ideally
// the bcrypt hash would be stored in the database along with the username.
var UserPassword = "p@ssw0rd"

func Run(database *sql.DB, listenAddress, certPath, keyPath, apiToken string) error {
	e := echo.New()

	// @NOTE CORS is only enabled like this for development
	// it would be configured with the below line in production
	// and properly have allowed origins, etc set.
	// middleware.CORSWithConfig(middleware.CORSConfig{})

	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	statikFS, err := fs.New()
	if err != nil {
		return fmt.Errorf("error initializing statik FS %w", err)
	}

	h := http.FileServer(statikFS)

	// New up the account store with a database connection
	metricStore := storage.NewPgMetricStore(database)
	accountStore := storage.NewPgAccountStore(database, metricStore)
	tokenStore := storage.NewTokenStore()
	accountUpdateStore := storage.NewAccountUpdateStore()

	// Hash the hard coded password and pass it in to be checked
	// @NOTE normally we'd hash the user's password on sign-up and store it in the DB
	hashedpw, err := hashUserPassword(UserPassword)
	if err != nil {
		return err
	}

	e.POST("/metrics", IngestMetricsHandler(apiToken, metricStore, accountUpdateStore))
	e.POST("/login", UserLoginHandler(tokenStore, UserEmail, hashedpw))

	// Protected routes
	g := e.Group("auth")
	g.Use(Authentication(tokenStore))
	g.POST("/logout", UserLogOutHandler(tokenStore))

	agrp := e.Group("accounts")
	agrp.Use(Authentication(tokenStore))
	agrp.POST("/:id/upgrade", AccountUpgradeHandler(accountStore))

	// @NOTE websockets cannot send headers, so this is unauthenticated but
	// will expect a message with a valid token to stay open on connect
	e.GET("/accounts/:id/updates", AccountUpdatesHandler(tokenStore, accountStore, accountUpdateStore))

	// the SPA route
	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", h)))

	return e.StartTLS(listenAddress, certPath, keyPath)
}
