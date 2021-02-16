package http

import (
	"crypto/subtle"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// As talked about in the design doc, I'm hard coding this. Ideally it would be stored
// in the database and looked up on request
const USER_EMAIL = "test@example.com"

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserLoginHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {

		loginRequest := UserLoginRequest{}

		if err := ctx.Bind(&loginRequest); err != nil {
			// Log out the error to the logger to capture the actual error
			log.Printf("error on user login %v\n", err)

			// Do not show the actual error to the user to allow them to see
			// anything about whether the email is valid/invalid, etc. to prevent anyone from
			// brute force looking for active accounts
			return ctx.JSON(http.StatusForbidden, HTTPError{"username or password is incorrect"})
		}

		// @NOTE there would be code here to go to the database and make sure that email exists as a valid account and has a password

		// check the email matches
		if subtle.ConstantTimeCompare([]byte(loginRequest.Email), []byte(USER_EMAIL)) == 0 {
			log.Printf("error user attempted to log in with invalid email\n")
			return ctx.JSON(http.StatusForbidden, HTTPError{"username or password is incorrect"})
		}

		cookie := createTokenCookie()

		ctx.SetCookie(cookie)

		// Normally I would set the domain from whatever site we're hosting on but for this case in development
		// I'll forego it
		// cookie.Domain = '.example.com'

		return ctx.NoContent(http.StatusNoContent)
	}
}

func UserLogOutHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := logUserOut(ctx)

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, HTTPError{"error logging out"})
		}

		ctx.SetCookie(cookie)

		return ctx.JSON(http.StatusNoContent, "logged out")
	}
}
