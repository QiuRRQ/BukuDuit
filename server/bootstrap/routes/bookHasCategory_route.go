package routes

import (
	api "bukuduit-go/server/handlers"

	"github.com/labstack/echo"
)

type BookHasCategoryRoute struct {
	RouteGroup *echo.Group
	Handler    api.Handler
}

func (route BookHasCategoryRoute) RegisterRoute() {
	handler := api.BookHasCategoryHandler{Handler: route.Handler}
	// jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	// route.RouteGroup.Use(jwtMiddleware.JWTWithConfig)
	route.RouteGroup.GET("/data", handler.Read)
	route.RouteGroup.GET("/:id", handler.ReadByID)
	route.RouteGroup.POST("", handler.Add)
	route.RouteGroup.PUT("/:id", handler.Edit)
	route.RouteGroup.DELETE("/:id", handler.Delete)
}
