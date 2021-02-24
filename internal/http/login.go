package http

import (
	"log"
	"net/http"

	"github.com/jesseobrien/fake-iot/internal/storage"
	"github.com/labstack/echo/v4"
)

const UserAccountId = ""

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoggedInResponse struct {
	Token     string `json:"access_token"`
	AccountId string `json:"account_id"`
}

func UserLoginHandler(tokenStore *storage.TokenStore, expectedEmail string, expectedPassword []byte) echo.HandlerFunc {
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

		// @NOTE In the interest of brevity I'm hardcoding the account ID here. It would be looked up when the user
		// logs in. I would send back the account ID and possibly the user ID as well to be able to make requests
		accountId := "47f3c307-6344-49e7-961c-ea200e950a89"

		return ctx.JSON(http.StatusOK, UserLoggedInResponse{token, accountId})
	}
}
