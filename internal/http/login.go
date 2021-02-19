package http

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserLoginHandler(expectedEmail string, expectedPassword []byte) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		loginRequest := UserLoginRequest{}

		if err := ctx.Bind(&loginRequest); err != nil {
			// Log out the error to the logger to capture the actual error
			log.Printf("error on user login %v\n", err)

			// Do not show the actual error to the user to allow them to see
			// anything about whether the email is valid/invalid, etc. to prevent anyone from
			// brute force looking for active accounts
			return echo.NewHTTPError(http.StatusUnauthorized, "something went wrong, try again or contact a site administrator")
		}

		// @NOTE there would be code here to go to the database and pull the user record instead of using the hard
		// coded values

		// If the user credentials are empty
		if loginRequest == (UserLoginRequest{}) {
			return echo.NewHTTPError(http.StatusUnauthorized, "you must provide a username and password")
		}

		if err := checkEmailAndPassword(loginRequest.Email, loginRequest.Password, expectedEmail, expectedPassword); err != nil {
			log.Printf("error user attempted to log in with invalid credentials: %v", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "username or password is incorrect")
		}

		cookie, err := createTokenCookie(loginRequest.Email)

		if err != nil {
			log.Printf("error generating cookie for user: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "something went wrong on our end, check with a site administrator")
		}

		ctx.SetCookie(cookie)

		return ctx.NoContent(http.StatusNoContent)
	}
}

func UserLogOutHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := expireLoginToken(ctx)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "error logging out")
		}

		ctx.SetCookie(cookie)

		return ctx.JSON(http.StatusNoContent, "logged out")
	}
}
