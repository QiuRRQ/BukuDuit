package bootstrap

import (
	"bukuduit-go/helpers/jwe"
	"bukuduit-go/helpers/jwt"
	"bukuduit-go/usecase"
	"database/sql"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Bootstrap struct {
	E                       *echo.Echo
	Db                      *sql.DB
	Redis                   *redis.Client
	UseCaseContract         usecase.UcContract
	Jwe                     jwe.Credential
	Validator               *validator.Validate
	Translator              ut.Translator
	JwtConfig               middleware.JWTConfig
	JwtCred                 jwt.JwtCredential
}

