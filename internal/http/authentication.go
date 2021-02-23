package http

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Store user tokens in memory.
// Note: Nominally these would be in a storage location like redis or something
// to be able to query for them across multiple services
var ErrInvalidAuthToken = errors.New("auth token is invalid")

func getBearerTokenFromHeader(ctx echo.Context) string {
	authHeader := ctx.Request().Header.Get("Authorization")

	token := strings.TrimPrefix(authHeader, "Bearer ")

	return token
}

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

func createUserToken(tokenStore TokenStore, email string) (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	newToken := hex.EncodeToString(token)

	tokenStore[newToken] = email

	return newToken, nil
}

func Authentication(tokenStore TokenStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			token := getBearerTokenFromHeader(ctx)

			if _, ok := tokenStore[token]; !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid auth token")
			}

			return next(ctx)
		}
	}
}

func expireLoginToken(ctx echo.Context, tokenStore TokenStore) {
	token := getBearerTokenFromHeader(ctx)
	// Delete the token from the registered tokens
	delete(tokenStore, token)
}
