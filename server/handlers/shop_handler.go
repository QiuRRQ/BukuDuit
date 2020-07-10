package handlers

import (
	"bukuduit-go/helpers/jwt"
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type ShopHandler struct {
	Handler
}

func (handler ShopHandler) BrowseByUser(ctx echo.Context) error {
	claim := ctx.Get("user").(*jwt.CustomClaims)
	uc := usecase.ShopUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseByUser(claim.Id)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler ShopHandler) ExportFile(ctx echo.Context) error {
	uc := usecase.ShopUseCase{UcContract: handler.UseCaseContract}
	ID := ctx.Param("id")
	status := ctx.QueryParam("lunas")
	name := ctx.QueryParam("name")
	res, err := uc.ExportToFile(ID, status, name)

	if err != nil {
		return err
	}

	return ctx.File(res)
}

func (handler ShopHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")
	status := ctx.QueryParam("lunas")
	name := ctx.QueryParam("name")
	uc := usecase.ShopUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Read(ID, status, name)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler ShopHandler) Edit(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(request.ShopRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.ShopUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(input, ID)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler ShopHandler) Add(ctx echo.Context) error {
	input := new(request.ShopRequest)
	claim := ctx.Get("user").(*jwt.CustomClaims)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.ShopUseCase{UcContract: handler.UseCaseContract}
	err := uc.Add(input, claim.Id)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler ShopHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.ShopUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}
