package routes

import (
	api "bukuduit-go/server/handlers"
	"github.com/labstack/echo"
)

type OtpRoutes struct {
	RouteGroup *echo.Group
	Handler    api.Handler
}

func (route OtpRoutes) RegisterRoute() {
	handler := api.OtpHandler{Handler:route.Handler}

	route.RouteGroup.POST("/request",handler.RequestOTP)
}
