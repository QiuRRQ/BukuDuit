package routes

import (
	api "bukuduit-go/server/handlers"
	middleware "bukuduit-go/server/middlewares"

	"github.com/labstack/echo"
)

type TransactionRoute struct {
	RouteGroup *echo.Group
	Handler    api.Handler
}

func (route TransactionRoute) RegisterRoute() {
	handler := api.TransactionHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	route.RouteGroup.Use(jwtMiddleware.JWTWithConfig)
	route.RouteGroup.GET("/:id", handler.Read)
	route.RouteGroup.DELETE("/:id", handler.Delete)
	route.RouteGroup.POST("/debt", handler.DebtPayment)
	// route.RouteGroup.GET("/DebtList", handler)
}
