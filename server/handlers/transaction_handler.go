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
	name := ctx.QueryParam("name")
	time := ctx.QueryParam("time")

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}

	res, err := uc.TransactionList(shopID, name, time)

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
	searching := ctx.QueryParam("search")
	name := ctx.QueryParam("name")
	amount := ctx.QueryParam("amount")
	date := ctx.QueryParam("date")
	startDate := ctx.QueryParam("start_date")
	endDate := ctx.QueryParam("end_date")
	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}

	res, err := uc.DebtReport(shopID, searching, name, amount, date, startDate, endDate)

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

func (handler TransactionHandler) DeleteDebt(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	err := uc.DeleteDebt(ID)

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
	err := uc.AddDebt(*input)

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

func (handler TransactionHandler) EditTransction(ctx echo.Context) error {
	input := new(request.TransactionRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	err := uc.EditTransaction(*input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler TransactionHandler) DeleteTrans(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	err := uc.DeleteTransactions(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}
