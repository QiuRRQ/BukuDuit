package handlers

import (
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type BorrowCardHandler struct {
	Handler
}

func (handler BorrowCardHandler) Read(ctx echo.Context) error {

	uc := usecase.BorrowCardUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Read()

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler BorrowCardHandler) LaporanPinjam(ctx echo.Context) error {

	uc := usecase.BorrowCardUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.LaporanPeminjaman()

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler BorrowCardHandler) LaporanPengembalian(ctx echo.Context) error {

	uc := usecase.BorrowCardUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.LaporanPengembalian()

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler BorrowCardHandler) ReadByID(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.BorrowCardUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadById(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler BorrowCardHandler) Edit(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(request.BorrowCardRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.BorrowCardUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(input, ID)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler BorrowCardHandler) Add(ctx echo.Context) error {
	input := new(request.BorrowCardRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.BorrowCardUseCase{UcContract: handler.UseCaseContract}
	err := uc.Add(input, nil)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler BorrowCardHandler) Pinjaman(ctx echo.Context) error {
	input := new(request.BorrowCardRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.BorrowCardUseCase{UcContract: handler.UseCaseContract}
	err := uc.AddPinjaman(input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler BorrowCardHandler) Kembalikan(ctx echo.Context) error {
	input := new(request.BorrowCardRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.BorrowCardUseCase{UcContract: handler.UseCaseContract}
	err := uc.AddPengembalian(input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler BorrowCardHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.BorrowCardUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}
