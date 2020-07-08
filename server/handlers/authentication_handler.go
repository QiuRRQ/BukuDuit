package handlers

import (
	"bukuduit-go/helpers/messages"
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type AuthenticationHandler struct {
	Handler
}

func (handler AuthenticationHandler) Register(ctx echo.Context) error {
	input := new(request.RegisterRequest) //i add a new filed called fullName since its requred on DB

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}
	uc := usecase.AuthenticationUseCase{UcContract: handler.UseCaseContract}
	err := uc.Register(input.MobilePhone, input.Pin, input.ShopName)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler AuthenticationHandler) Login(ctx echo.Context) error {
	input := new(request.LoginRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.AuthenticationUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Login(input.MobilePhone, input.PIN)
	if err != nil {
		return handler.SendResponseUnauthorized(ctx, err)
	}

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler AuthenticationHandler) GenerateTokenByOtp(ctx echo.Context) error {
	input := new(request.OtpSubmitRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}
	uc := usecase.AuthenticationUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.GenerateTokenByOtp("otp"+input.MobilePhone, input.Otp)

	if err != nil && err.Error() == messages.OtpNotMatch {
		return handler.SendResponseUnauthorized(ctx, err)
	}

	return handler.SendResponse(ctx, res, nil, err)
}
