package handlers

import (
	"bukuduit-go/helpers/DtoResponse"
	"bukuduit-go/helpers/jwe"
	"bukuduit-go/helpers/str"
	"bukuduit-go/usecase"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strings"
)

type Handler struct {
	E               *echo.Echo
	UseCaseContract *usecase.UcContract
	Jwe             jwe.Credential
	Db              *sql.DB
	Validate        *validator.Validate
	Translator      ut.Translator
	JwtConfig       middleware.JWTConfig
}

func (h Handler) SendResponse(ctx echo.Context, data interface{}, pagination interface{}, err error) error {
	response := DtoResponse.SuccessResponse(data, pagination)
	if err != nil {
		response = DtoResponse.ErrorResponse(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(response.StatusCode, response.Body)
}

func (h Handler) SendResponseBadRequest(ctx echo.Context, statusCode int, err interface{}) error {
	response := DtoResponse.ErrorResponse(statusCode, err)

	return ctx.JSON(response.StatusCode, response.Body)
}

func (h Handler) SendResponseErrorValidation(ctx echo.Context, error validator.ValidationErrors) error {
	errorMessages := h.ExtractErrorValidationMessages(error)

	return h.SendResponseBadRequest(ctx, http.StatusBadRequest, errorMessages)
}

func (h Handler) SendResponseUnauthorized(ctx echo.Context, err error) error {
	response := DtoResponse.ErrorResponse(http.StatusUnauthorized, err.Error())

	return ctx.JSON(response.StatusCode, response.Body)
}

func (h Handler) ResponseBadRequest(error string) viewmodel.ResponseVm {
	responseVm := viewmodel.ResponseVm{
		Body: viewmodel.RespBodyVm{
			Message:    error,
			DataVm:     nil,
			Pagination: nil,
		},
		StatusCode: http.StatusBadRequest,
	}

	return responseVm
}

func (h Handler) ResponseValidationError(error validator.ValidationErrors) viewmodel.ResponseVm {
	errorMessage := map[string][]string{}
	errorTranslation := error.Translate(h.Translator)

	for _, err := range error {
		errKey := str.Underscore(err.StructField())
		errorMessage[errKey] = append(
			errorMessage[errKey],
			strings.Replace(errorTranslation[err.Namespace()], err.StructField(), err.StructField(), -1),
		)
	}

	response := viewmodel.ResponseVm{
		Body: viewmodel.RespBodyVm{
			Message:    errorMessage,
			DataVm:     nil,
			Pagination: nil,
		},
		StatusCode: http.StatusBadRequest,
	}

	return response
}

func (h Handler) ExtractErrorValidationMessages(error validator.ValidationErrors) map[string][]string {
	errorMessage := map[string][]string{}
	errorTranslation := error.Translate(h.Translator)

	for _, err := range error {
		errKey := str.Underscore(err.StructField())
		errorMessage[errKey] = append(
			errorMessage[errKey],
			strings.Replace(errorTranslation[err.Namespace()], err.StructField(), err.StructField(), -1),
		)
	}

	return errorMessage
}