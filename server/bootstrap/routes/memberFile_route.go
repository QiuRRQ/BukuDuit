package routes

import (
	api "bukuduit-go/server/handlers"

	"github.com/labstack/echo"
)

type MemberFileRoute struct {
	RouteGroup *echo.Group
	Handler    api.Handler
}

func (route MemberFileRoute) RegisterRoute() {
	handler := api.MemberFileHandler{Handler: route.Handler}
	// jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	// route.RouteGroup.Use(jwtMiddleware.JWTWithConfig)
	route.RouteGroup.POST("", handler.Add)
}
