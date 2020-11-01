package handlers

import (
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type MembersHandler struct {
	Handler
}

func (handler MembersHandler) Read(ctx echo.Context) error {

	uc := usecase.MembersUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Read()

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler MembersHandler) ReadByID(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.MembersUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadById(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler MembersHandler) Edit(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(request.MemberRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.MembersUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(input, ID)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler MembersHandler) Add(ctx echo.Context) error {
	input := new(request.MemberRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.MembersUseCase{UcContract: handler.UseCaseContract}
	err := uc.Add(input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler MembersHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.MembersUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}
