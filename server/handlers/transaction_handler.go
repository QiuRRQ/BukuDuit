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

// func (handler TransactionHandler) BrowseByCustomer(ctx echo.Context) error {
// 	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}

// 	return handler.SendResponse(ctx, res, nil, err)
// }

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

func (handler TransactionHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}

//untuk pembayaran hutang
func (handler TransactionHandler) DebtPayment(ctx echo.Context) error {
	input := new(request.UtangRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	err := uc.DebtPayment(input.Reference_id, input.DebtType, input.Shop_id, input.Amount)

	return handler.SendResponse(ctx, nil, nil, err)
}
