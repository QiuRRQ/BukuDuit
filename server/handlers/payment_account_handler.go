package handlers

import (
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type PaymentAccountHandler struct {
	Handler
}

func (handler PaymentAccountHandler) BrowseByShop(ctx echo.Context) error {
	input := ctx.QueryParam("shopid")
	uc := usecase.PaymentAccountUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseByShop(input)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler PaymentAccountHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")
	status := ctx.QueryParam("lunas")
	uc := usecase.PaymentAccountUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Read(ID, status)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler PaymentAccountHandler) Edit(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(request.PaymentAccountRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.PaymentAccountUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(input, ID)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler PaymentAccountHandler) Add(ctx echo.Context) error {
	input := new(request.PaymentAccountRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.PaymentAccountUseCase{UcContract: handler.UseCaseContract}
	err := uc.Add(input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler PaymentAccountHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.PaymentAccountUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}
