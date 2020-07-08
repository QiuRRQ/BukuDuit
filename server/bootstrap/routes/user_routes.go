package routes

import (
	api "bukuduit-go/server/handlers"

	"github.com/labstack/echo"
)

type UserRoute struct {
	RouteGroup *echo.Group
	Handler    api.Handler
}

func (route UserRoute) RegisterRoute() {
	handler := api.UserHandler{Handler: route.Handler}

	route.RouteGroup.GET("/:id", handler.ReadAccountDetail)
	route.RouteGroup.POST("/:id", handler.Edit)
	route.RouteGroup.POST("/forgot", handler.ForgotPin)
}
