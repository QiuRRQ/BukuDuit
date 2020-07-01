package handlers

import (
	"bukuduit-go/usecase"

	"github.com/labstack/echo"
)

type UserHandler struct {
	Handler
}

func (handler UserHandler) ReadAccountDetail(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.UserUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.MyAccount(ID)

	return handler.SendResponse(ctx, res, nil, err)
}
