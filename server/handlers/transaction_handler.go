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

func (handler TransactionHandler) TransactionListWeeks(ctx echo.Context) error {
	shopID := ctx.QueryParam("shopid")
	searching := ctx.QueryParam("search")
	name := ctx.QueryParam("name")
	amount := ctx.QueryParam("amount")
	date := ctx.QueryParam("date")
	timeGroup := ctx.QueryParam("time")
	startDate := ctx.QueryParam("start_date")
	endDate := ctx.QueryParam("end_date")

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}

	res, err := uc.TransactionListByWeeks(shopID, searching, name, amount, date, timeGroup, startDate, endDate)

	return handler.SendResponse(ctx, res, nil, err)
}

//list transaksi by months
func (handler TransactionHandler) TransactionListMonth(ctx echo.Context) error {
	shopID := ctx.QueryParam("shopid")
	searching := ctx.QueryParam("search")
	name := ctx.QueryParam("name")
	amount := ctx.QueryParam("amount")
	date := ctx.QueryParam("date")
	startDate := ctx.QueryParam("start_date")
	endDate := ctx.QueryParam("end_date")

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}

	res, err := uc.TransactionListMonth(shopID, searching, name, amount, date, startDate, endDate)

	return handler.SendResponse(ctx, res, nil, err)
}

//list transaksi by days
func (handler TransactionHandler) TransactionList(ctx echo.Context) error {
	shopID := ctx.QueryParam("shopid")
	searching := ctx.QueryParam("search")
	name := ctx.QueryParam("name")
	amount := ctx.QueryParam("amount")
	date := ctx.QueryParam("date")
	startDate := ctx.QueryParam("start_date")
	endDate := ctx.QueryParam("end_date")

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}

	res, err := uc.TransactionList(shopID, searching, name, amount, date, "", startDate, endDate)

	return handler.SendResponse(ctx, res, nil, err)
}
func (handler TransactionHandler) BrowseByCustomer(ctx echo.Context) error {
	customerId := ctx.QueryParam("customerId")

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseByCustomer(customerId)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler TransactionHandler) TransactionReport(ctx echo.Context) error {
	shopID := ctx.QueryParam("shopid")
	searching := ctx.QueryParam("search")
	name := ctx.QueryParam("name")
	amount := ctx.QueryParam("amount")
	date := ctx.QueryParam("date")
	startDate := ctx.QueryParam("start_date")
	endDate := ctx.QueryParam("end_date")
	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}

	res, err := uc.TransactionReport(shopID, searching, name, amount, date, startDate, endDate)
	if err != nil {
		return err
	}

	return handler.SendResponse(ctx, res, nil, err)
}

//export excel untuk laporan transaksi
func (handler TransactionHandler) TransactionReportExportFile(ctx echo.Context) error {
	shopID := ctx.QueryParam("shopid")
	searching := ctx.QueryParam("search")
	name := ctx.QueryParam("name")
	amount := ctx.QueryParam("amount")
	date := ctx.QueryParam("date")
	startDate := ctx.QueryParam("start_date")
	endDate := ctx.QueryParam("end_date")
	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}

	res, err := uc.TransactionReportExportFile(shopID, searching, name, amount, date, startDate, endDate)
	if err != nil {
		return err
	}

	return ctx.File(res)
}

//export excel untuk utang detail
func (handler TransactionHandler) DebtDetailExportFile(ctx echo.Context) error {
	customerID := ctx.QueryParam("customerId")

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}

	res, err := uc.DebtDetailExportFile(customerID)

	if err != nil {
		return err
	}

	return ctx.File(res)
}

//export excel untuk laporan hutang
func (handler TransactionHandler) DebtReportExportFile(ctx echo.Context) error {
	shopID := ctx.QueryParam("shopid")
	searching := ctx.QueryParam("search")
	name := ctx.QueryParam("name")
	amount := ctx.QueryParam("amount")
	date := ctx.QueryParam("date")
	startDate := ctx.QueryParam("start_date")
	endDate := ctx.QueryParam("end_date")
	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}

	res, err := uc.DebtReportExportFile(shopID, searching, name, amount, date, startDate, endDate)

	if err != nil {
		return err
	}

	return ctx.File(res)
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
