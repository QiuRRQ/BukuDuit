package handlers

import (
	"bukuduit-go/usecase"

	"github.com/labstack/echo"
)

type MemberFileHandler struct {
	Handler
}

func (handler MemberFileHandler) Add(ctx echo.Context) error {

	uc := usecase.MemberFileUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Add(ctx)

	return handler.SendResponse(ctx, res, nil, err)
}
