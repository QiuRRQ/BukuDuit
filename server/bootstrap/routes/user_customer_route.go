package routes

import (
	api "bukuduit-go/server/handlers"
	middleware "bukuduit-go/server/middlewares"
	"github.com/labstack/echo"
)

type UserCustomerRoute struct {
	RouteGroup *echo.Group
	Handler    api.Handler
}

func (route UserCustomerRoute) RegisterRoute() {
	handler := api.UserCustomerHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	route.RouteGroup.Use(jwtMiddleware.JWTWithConfig)
	route.RouteGroup.GET("",handler.BrowseByShop)
	route.RouteGroup.GET("/:id",handler.Read)
	route.RouteGroup.POST("",handler.Add)
	route.RouteGroup.DELETE("/:id",handler.Delete)
}
