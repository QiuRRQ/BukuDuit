package handlers

import (
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
)

type AuthenticationHandler struct {
	Handler
}

func (handler AuthenticationHandler) Register(ctx echo.Context) error {
	input := new(request.RegisterRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}
	uc := usecase.AuthenticationUseCase{UcContract: handler.UseCaseContract}
	err := uc.Register(input.MobilePhone, input.Pin)

	return handler.SendResponse(ctx, nil, nil, err)
}
