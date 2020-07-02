package handlers

import (
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type TransactionHandler struct {
	Handler
}

func (handler TransactionHandler) TransactionList(ctx echo.Context) error {
	shopID := ctx.QueryParam("shopid")

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}

	res, err := uc.TransactionList(shopID)

	return handler.SendResponse(ctx, res, nil, err)
}
func (handler TransactionHandler) BrowseByCustomer(ctx echo.Context) error {
	customerId := ctx.QueryParam("customerId")

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseByCustomer(customerId)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler TransactionHandler) BrowseByShop(ctx echo.Context) error {
	shopID := ctx.QueryParam("shopid")

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}

	res, err := uc.BrowseByShop(shopID)

	return handler.SendResponse(ctx, res, nil, err)
}
func (handler TransactionHandler) BrowseUser(ctx echo.Context) error {
	ID := ctx.QueryParam("userID")
	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseByCustomer(ID)

	return handler.SendResponse(ctx, res, nil, err)
}
func (handler TransactionHandler) Read(ctx echo.Context) error {
	ID := ctx.QueryParam("id")
	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Read(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler TransactionHandler) Edit(ctx echo.Context) error {
	input := new(request.TransactionRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	err := uc.EditDebt(*input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler TransactionHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}

//untuk pembayaran hutang
func (handler TransactionHandler) DebtPayment(ctx echo.Context) error {
	input := new(request.TransactionRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	err := uc.DebtPayment(*input)

	return handler.SendResponse(ctx, nil, nil, err)
}

//untuk add transaksi
func (handler TransactionHandler) AddTransaction(ctx echo.Context) error {
	input := new(request.TransactionRequest) //id using users ID not UserCustomerID

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	err := uc.AddTransaksi(*input)

	return handler.SendResponse(ctx, nil, nil, err)
}
