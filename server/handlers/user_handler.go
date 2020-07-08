package handlers

import (
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type UserHandler struct {
	Handler
}

func (handler UserHandler) ReadAccountDetail(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.UserUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.MyAccount(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler UserHandler) Edit(ctx echo.Context) error {
	input := new(request.UserRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.UserUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(*input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler UserHandler) ForgotPin(ctx echo.Context) error {
	input := new(request.UserRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.UserUseCase{UcContract: handler.UseCaseContract}
	err := uc.ForgotMyPin(*input)

	return handler.SendResponse(ctx, nil, nil, err)
}
