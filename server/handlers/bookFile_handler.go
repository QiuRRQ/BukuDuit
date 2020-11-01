package handlers

import (
	"bukuduit-go/usecase"

	"github.com/labstack/echo"
)

type BookFileHandler struct {
	Handler
}

func (handler BookFileHandler) Add(ctx echo.Context) error {

	uc := usecase.BookFileUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Add(ctx, "")

	return handler.SendResponse(ctx, res, nil, err)
}
