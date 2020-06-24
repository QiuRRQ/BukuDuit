package handlers

import (
	"bukuduit-go/helpers/jwt"
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type TransactionHandler struct {
	Handler
}

func (handler TransactionHandler) BrowseByCustomer(ctx echo.Context) error {
	fmt.Printf("%+v\n", claim)
	fmt.Println("coba")
	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseByCustomer(claim.Id)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler TransactionHandler) BrowseByDebtType(ctx echo.Context) error{

	return handler.SendResponse(ctx, res, nil, err)
}
func (handler TransactionHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Read(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler TransactionHandler) Edit(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(request.TransactionRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(input, ID)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler TransactionHandler) Add(ctx echo.Context) error {
	input := new(request.TransactionRequest)
	claim := ctx.Get("customer").(*jwt.CustomClaims)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	err := uc.Add(input, claim.Id)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler TransactionHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}

//untuk pembayaran hutang
func (handler TransactionHandler) DebtPayment(ctx echo.Context) error{
	input := new(request.UtangRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	err := uc.DebtPayment(input.CustomerID, input.DebtType, input.UserCustomerDebt, input.Amount)

	return handler.SendResponse(ctx, nil, nil, err)
}