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
	route.RouteGroup.GET("/list/:userID", handler.BrowseUser) // list transaksi done
	route.RouteGroup.DELETE("/:id", handler.Delete)
	route.RouteGroup.POST("/transaction", handler.AddTransaction) //done
	route.RouteGroup.GET("/details/:id", handler.Read)            //done ini jadi satu utang dan transaksi
	route.RouteGroup.POST("/debt", handler.DebtPayment)
	route.RouteGroup.GET("/debt", handler.BrowseByCustomer) //list detail utang
	route.RouteGroup.GET("/debtreport", handler.BrowseByShop)
}
