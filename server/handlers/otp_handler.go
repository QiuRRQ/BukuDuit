package handlers

import (
	"net/http"
	"bukuduit-go/server/requests"
	"bukuduit-go/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type OtpHandler struct {
	Handler
}

func (handler OtpHandler) RequestOTP(ctx echo.Context) error {
	input := new(request.OtpRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	otpUc := usecase.OtpUseCase{UcContract: handler.UseCaseContract}
	otpUc.SetXRequestID(ctx)
	_, err := otpUc.RequestOtp(input.MobilePhone,input.Action)

	return handler.SendResponse(ctx, nil, nil, err)
}