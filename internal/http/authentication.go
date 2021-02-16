package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const AUTH_COOKIE_NAME = "user_token"

// Store user tokens in memory.
// Note: Nominally these would be in a storage location like redis or something
// to be able to query for them across multiple services
var USER_TOKENS = map[string]string{}

type ErrInvalidAuthToken struct {
}

func (e ErrInvalidAuthToken) Error() string {
	return "auth token is invalid"
}

// Create a login cookie with a new token
func createTokenCookie() *http.Cookie {
	newToken := "abc"

	cookie := new(http.Cookie)
	cookie.Name = AUTH_COOKIE_NAME
	cookie.Value = newToken
	cookie.Expires = time.Now().Add(30 * time.Hour)

	// Ensure the cookie only works over HTTPS
	cookie.Secure = true

	return cookie
}

func checkLoginCookie(ctx echo.Context) error {
	cookie, err := ctx.Cookie(AUTH_COOKIE_NAME)

	if err != nil {
		return fmt.Errorf("error getting cookie `%s`, %w", AUTH_COOKIE_NAME, err)
	}

	// find the cookie value in the map
	_, ok := USER_TOKENS[cookie.Value]
	if !ok {
		return ErrInvalidAuthToken{}
	}

	return nil
}

func logUserOut(ctx echo.Context) (*http.Cookie, error) {
	cookie, err := ctx.Cookie(AUTH_COOKIE_NAME)
	if err != nil {
		return nil, err
	}

	// Set the max age of the cookie to 0, meaning it will expire immediately
	cookie.MaxAge = -1
	cookie.Expires = time.Now().Add(-24 * time.Hour)

	return cookie, nil
}
