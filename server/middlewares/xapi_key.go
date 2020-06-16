package middleware

import (
	"bukuduit-go/helpers/messages"
	"bukuduit-go/server/handlers"
	"bukuduit-go/usecase"
	"errors"
	"github.com/labstack/echo"
	"time"
)

type XApiKey struct {
	*usecase.UcContract
}

func (xApiKey XApiKey) VerifyXApiKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		plainText := "legoasauction"+time.Now().UTC().Format("02-Jan-2006")+"key"
		apiHandler := handlers.Handler{UseCaseContract: xApiKey.UcContract}

		authHeader := ctx.Request().Header.Get("x-api-key")
		if authHeader == "" {
			return apiHandler.SendResponseUnauthorized(ctx, errors.New(messages.AuthHeaderNotPresent))
		}

		res,err := xApiKey.AesFront.Decrypt(authHeader)
		if res != plainText {
			return apiHandler.SendResponseUnauthorized(ctx, errors.New(messages.Unauthorized))
		}

		return next(ctx)
	}
}
