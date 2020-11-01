package routes

import (
	api "bukuduit-go/server/handlers"

	"github.com/labstack/echo"
)

type BooksRoute struct {
	RouteGroup *echo.Group
	Handler    api.Handler
}

func (route BooksRoute) RegisterRoute() {
	handler := api.BooksHandler{Handler: route.Handler}

	route.RouteGroup.GET("/data", handler.Read)
	route.RouteGroup.GET("/:id", handler.ReadByID)
	route.RouteGroup.POST("", handler.Add)
	route.RouteGroup.PUT("/:id", handler.Edit)
	route.RouteGroup.PUT("/stock/:id", handler.EditStock)
	route.RouteGroup.DELETE("/:id", handler.Delete)
}
