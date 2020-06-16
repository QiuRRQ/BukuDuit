package bootstrap

import (
	api "bukuduit-go/server/handlers"
	"github.com/labstack/echo"
	"net/http"
)

func(boot *Bootstrap) RegisterRouters(){
	_ = api.Handler{
		E:               boot.E,
		Db:              boot.Db,
		UseCaseContract: &boot.UseCaseContract,
		Jwe:             boot.Jwe,
		Validate:        boot.Validator,
		Translator:      boot.Translator,
		JwtConfig:       boot.JwtConfig,
	}

	boot.E.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "Work")
	})
}