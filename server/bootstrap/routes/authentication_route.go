package routes

import (
	api "bukuduit-go/server/handlers"

	"github.com/labstack/echo"
)

type AuthenticationRoute struct {
	RouteGroup *echo.Group
	Handler    api.Handler
}

//register route
func (route AuthenticationRoute) RegisterRoute() {
	handler := api.AuthenticationHandler{Handler: route.Handler}

	route.RouteGroup.POST("/register", handler.Register)
	route.RouteGroup.POST("/login", handler.Login)
	route.RouteGroup.POST("/phone_number", handler.PhoneCheck)
	route.RouteGroup.POST("/by-otp", handler.GenerateTokenByOtp)

}
