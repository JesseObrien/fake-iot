package http

import (
	"crypto/rand"
	"crypto/subtle"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const AuthCookieName = "user_token"

// Store user tokens in memory.
// Note: Nominally these would be in a storage location like redis or something
// to be able to query for them across multiple services
var UserTokens = map[string]string{}

var ErrInvalidAuthToken = errors.New("auth token is invalid")

func checkEmailAndPassword(email, password, expectedEmail string, hashedPassword []byte) error {
	if subtle.ConstantTimeCompare([]byte(expectedEmail), []byte(email)) == 0 {
		return errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		return err
	}

	return nil
}

func hashUserPassword(password string) ([]byte, error) {
	hashedpw, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, fmt.Errorf("could not hash hard coded password %w", err)
	}

	return hashedpw, nil
}

// Create a login cookie with a new token
func createTokenCookie(email string) (*http.Cookie, error) {

	token := make([]byte, 26)
	_, err := rand.Read(token)
	if err != nil {
		return nil, err
	}

	newToken := fmt.Sprintf("%x", token)

	// Keep track of user tokens in a map so we can revoke them if need be
	// @NOTE this is not a good or secure way of holding the tokens. I would probably
	// store them in the DB user's table and be able to look them up/revoke them from there.
	UserTokens[newToken] = email

	cookie := new(http.Cookie)
	cookie.Name = AuthCookieName
	cookie.Value = newToken
	cookie.Expires = time.Now().Add(30 * time.Hour)

	cookie.SameSite = http.SameSiteStrictMode

	// Ensure the cookie only works over HTTPS
	cookie.Secure = true

	return cookie, nil
}

func checkLoginCookie(ctx echo.Context) error {
	cookie, err := ctx.Cookie(AuthCookieName)

	if err != nil {
		return fmt.Errorf("error getting cookie `%s`, %w", AuthCookieName, err)
	}

	// find the cookie value in the map
	_, ok := UserTokens[cookie.Value]
	if !ok {
		return ErrInvalidAuthToken
	}

	return nil
}

func expireLoginToken(ctx echo.Context) (*http.Cookie, error) {
	cookie, err := ctx.Cookie(AuthCookieName)
	if err != nil {
		return nil, err
	}

	// Delete the token from the registered tokens
	delete(UserTokens, cookie.Value)

	// Set the max age of the cookie to 0, meaning it will expire immediately
	cookie.MaxAge = -1
	cookie.Expires = time.Now().Add(-24 * time.Hour)

	return cookie, nil
}
