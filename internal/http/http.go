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
const USER_EMAIL = "test@example.com"

// As in the design doc, I'm hard coding the password and will hash it on startup. Ideally
// the bcrypt hash would be stored in the database along with the username.
var USER_PASSWORD = "p@ssw0rd"

func Run(database *sql.DB, listenAddress, certPath, keyPath, apiToken string) error {
	e := echo.New()

	e.Use(middleware.CORS())
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

	// Hash the hard coded password and pass it in to be checked
	// @NOTE normally we'd hash the user's password on sign-up and store it in the DB
	hashedpw, err := hashUserPassword(USER_PASSWORD)
	if err != nil {
		return err
	}

	e.POST("/metrics", IngestMetricsHandler(apiToken, accountStore))
	e.POST("/login", UserLoginHandler(USER_EMAIL, hashedpw))
	e.POST("/logout", UserLogOutHandler())

	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", h)))

	return e.StartTLS(listenAddress, certPath, keyPath)
}
