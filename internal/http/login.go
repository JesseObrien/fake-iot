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

type UserLoggedInResponse struct {
	Token string `json:"access_token"`
}

func UserLoginHandler(tokenStore TokenStore, expectedEmail string, expectedPassword []byte) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		loginRequest := UserLoginRequest{}

		if err := ctx.Bind(&loginRequest); err != nil {
			// Log out the error to the logger to capture the actual error
			log.Printf("error on user login %v", err)

			// Do not show the actual error to the user to allow them to see
			// anything about whether the email is valid/invalid, etc. to prevent anyone from
			// brute force looking for active accounts
			return echo.NewHTTPError(http.StatusUnauthorized, "something went wrong, try again or contact a site administrator")
		}

		// @NOTE there would be code here to go to the database and pull the user record instead of using the hard
		// coded values

		if err := checkEmailAndPassword(loginRequest.Email, loginRequest.Password, expectedEmail, expectedPassword); err != nil {
			log.Printf("error user attempted to log in with invalid credentials: %v", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "username or password is incorrect")
		}

		token, err := createUserToken(tokenStore, loginRequest.Email)

		if err != nil {
			log.Printf("error generating token for user: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "something went wrong on our end, check with a site administrator")
		}

		return ctx.JSON(http.StatusOK, UserLoggedInResponse{token})
	}
}
