package handlers

import (
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
)

type BusinessCardHandler struct {
	Handler
}

func (handler BusinessCardHandler) BrowseByUser(ctx echo.Context) error {
	userID := ctx.QueryParam("userId")
	uc := usecase.BusinessCardUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseByUser(userID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler BusinessCardHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.BusinessCardUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Read(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler BusinessCardHandler) Edit(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(request.BusinessCardRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.BusinessCardUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(input, ID)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler BusinessCardHandler) Add(ctx echo.Context) error {
	input := new(request.BusinessCardRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.BusinessCardUseCase{UcContract: handler.UseCaseContract}
	err := uc.Add(input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler BusinessCardHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.BusinessCardUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}
