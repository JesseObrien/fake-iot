package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/jesseobrien/fake-iot/internal/storage"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

type SocketOperation string

const (
	SocketError                    SocketOperation = "error"
	AuthorizationFailure           SocketOperation = "authorization_failure"
	AccountUpdatesSubscribeRequest SocketOperation = "account_updates_subscribe"
	AccountInfoResponse            SocketOperation = "account_info_response"
	AccountMetricsUpdated          SocketOperation = "account_metrics_updated"
)

type SocketMessage struct {
	Operation SocketOperation `json:"operation"`
	Token     string          `json:"token,omitempty"`
	Data      string          `json:"data,omitempty"`
}

type AccountUpdatesRequest struct {
	AccountId string `path:"id"`
}

func AccountUpdatesHandler(tokenStore *storage.TokenStore, accountStore storage.AccountStore, accountUpdateStore *storage.AccountUpdateStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		updateRequest := AccountUpdatesRequest{}
		if err := echo.PathParamsBinder(ctx).MustString("id", &updateRequest.AccountId).BindError(); err != nil {
			ctx.Logger().Error(err)
			return err
		}

		socket, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
		if err != nil {
			return err
		}

		defer socket.Close()

		for {
			msg := SocketMessage{}

			// Block until a message is sent across the socket
			if err := socket.ReadJSON(&msg); err != nil {
				ctx.Logger().Error(err)
				break
			}

			// Check that a valid token has been sent in the message payload
			if err := checkMessageToken(ctx, msg, tokenStore, socket); err != nil {
				ctx.Logger().Error(err)
				break
			}

			// React to the different message types
			switch msg.Operation {
			case AccountUpdatesSubscribeRequest:
				if err := handleAccountUpdatesSubscribeRequest(ctx.Request().Context(), accountStore, updateRequest.AccountId, socket); err != nil {
					ctx.Logger().Error(err)
					break
				}

				subscription := accountUpdateStore.Subscribe(updateRequest.AccountId)

				// @TODO subscribe to the channel for the account to send updates for the account ID
				go handleAccountUpdatesSubscription(ctx, subscription, socket)

			default:
				if err := socket.WriteJSON(SocketMessage{Operation: SocketError, Data: fmt.Sprintf("unregistered operation: %s", msg.Operation)}); err != nil {
					ctx.Logger().Error(err)
				}
			}
		}
		return nil
	}
}

func checkMessageToken(ctx echo.Context, msg SocketMessage, tokenStore *storage.TokenStore, socket *websocket.Conn) error {
	token := strings.TrimPrefix(msg.Token, "Bearer ")
	if tokenStore.IsValid(token) {
		return nil
	}

	// Let the front end know there's an authorization problem
	if err := socket.WriteJSON(SocketMessage{Operation: AuthorizationFailure, Data: "invalid auth token"}); err != nil {
		return err
	}

	return errors.New("websocket message sent with invalid authorization token")
}

func handleAccountUpdatesSubscription(ctx echo.Context, subscription *storage.AccountUpdateSubscription, socket *websocket.Conn) {
	defer subscription.Unsubscribe()

	for {
		accountUpdate := <-subscription.Updates

		jsonUpdate, _ := json.Marshal(accountUpdate)

		if err := socket.WriteJSON(SocketMessage{Operation: AccountMetricsUpdated, Data: string(jsonUpdate)}); err != nil {
			ctx.Logger().Error(err)
		}
	}
}

func handleAccountUpdatesSubscribeRequest(ctx context.Context, accountStore storage.AccountStore, accountId string, socket *websocket.Conn) error {
	accountInfo, err := accountStore.GetAccountInfo(ctx, accountId)
	if err != nil {
		// Let the front-end know there's been an error
		// @NOTE this is a generic error. Ideally I'd abstract out having Request/Response objects with
		// error states on them so the request would know if it's successful/failed.
		if err := socket.WriteJSON(SocketMessage{Operation: SocketError, Data: err.Error()}); err != nil {
			return err
		}
		return fmt.Errorf("could not get account info %w", err)
	}

	jsonInfo, _ := json.Marshal(accountInfo)
	if err := socket.WriteJSON(SocketMessage{Operation: AccountInfoResponse, Data: string(jsonInfo)}); err != nil {
		return err
	}

	return nil
}
