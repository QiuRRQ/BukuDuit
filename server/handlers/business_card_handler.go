package handlers

import (
	"bukuduit-go/helpers/jwt"
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type BusinessCardHandler struct {
	Handler
}

func (handler BusinessCardHandler) BrowseByUser(ctx echo.Context) error {
	claim := ctx.Get("user").(*jwt.CustomClaims)
	uc := usecase.BusinessCardUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseByUser(claim.Id)

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
	claim := ctx.Get("user").(*jwt.CustomClaims)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.BusinessCardUseCase{UcContract: handler.UseCaseContract}
	err := uc.Add(input, claim.Id)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler BusinessCardHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.BusinessCardUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}
