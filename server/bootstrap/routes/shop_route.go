package routes

import (
	api "bukuduit-go/server/handlers"
	middleware "bukuduit-go/server/middlewares"
	"github.com/labstack/echo"
)

type ShopRoute struct {
	RouteGroup *echo.Group
	Handler api.Handler
}

func (route ShopRoute) RegisterRoute(){
	handler := api.BusinessCardHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	route.RouteGroup.Use(jwtMiddleware.JWTWithConfig)
	route.RouteGroup.GET("",handler.BrowseByUser)
	route.RouteGroup.GET("/:id",handler.Read)
	route.RouteGroup.POST("",handler.Add)
	route.RouteGroup.PUT("/:id",handler.Edit)
	route.RouteGroup.DELETE("/:id",handler.Delete)
}
