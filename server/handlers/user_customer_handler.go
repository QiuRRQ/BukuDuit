package handlers

import (
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type UserCustomerHandler struct {
	Handler
}

func (handler UserCustomerHandler) BrowseByShop(ctx echo.Context) error {
	shopId := ctx.QueryParam("shopId")

	uc := usecase.UserCustomerUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseByShop(shopId)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler UserCustomerHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.UserCustomerUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Read(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler UserCustomerHandler) Add(ctx echo.Context) error {
	input := new(request.UserCustomerRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.UserCustomerUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Add(input)

	if res != "" {
		return handler.SendResponse(ctx, res, nil, nil)
	}

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler UserCustomerHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.UserCustomerUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}
