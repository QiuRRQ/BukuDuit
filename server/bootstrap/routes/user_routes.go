package routes

import (
	api "bukuduit-go/server/handlers"
	middleware "bukuduit-go/server/middlewares"

	"github.com/labstack/echo"
)

type UserRoute struct {
	RouteGroup *echo.Group
	Handler    api.Handler
}

func (route UserRoute) RegisterRoute() {
	handler := api.UserHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	route.RouteGroup.Use(jwtMiddleware.JWTWithConfig)
	route.RouteGroup.GET("/:id", handler.ReadAccountDetail)
}
