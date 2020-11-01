package routes

import (
	api "bukuduit-go/server/handlers"

	"github.com/labstack/echo"
)

type BarcodeRoute struct {
	RouteGroup *echo.Group
	Handler    api.Handler
}

func (route BarcodeRoute) RegisterRoute() {
	handler := api.BarcodeHandler{Handler: route.Handler}
	// jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	// route.RouteGroup.Use(jwtMiddleware.JWTWithConfig)
	route.RouteGroup.GET("/data", handler.Read)
	route.RouteGroup.GET("/:id", handler.ReadByID)
	route.RouteGroup.POST("", handler.Add)
	route.RouteGroup.DELETE("/:id", handler.Delete)
}
