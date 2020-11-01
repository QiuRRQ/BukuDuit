package handlers

import (
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type BooksHandler struct {
	Handler
}

func (handler BooksHandler) Read(ctx echo.Context) error {

	uc := usecase.BooksUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Read()

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler BooksHandler) ReadByID(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.BooksUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadById(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler BooksHandler) EditStock(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(request.BookRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.BooksUseCase{UcContract: handler.UseCaseContract}
	err := uc.EditStockRequest(input, ID)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler BooksHandler) Edit(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(request.BookRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.BooksUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(input, ID)

	return handler.SendResponse(ctx, nil, nil, err)
}

//ini untuk menambah buku,
func (handler BooksHandler) Add(ctx echo.Context) error {
	input := new(request.BookRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.BooksUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Add(input, nil)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler BooksHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.BooksUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}
