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
	route.RouteGroup.GET("/list", handler.TransactionList)
	route.RouteGroup.DELETE("/:id", handler.Delete)
	route.RouteGroup.POST("/transaction", handler.AddTransaction)     //done
	route.RouteGroup.POST("/transactionedit", handler.AddTransaction) //change to edit
	route.RouteGroup.GET("/details/:id", handler.Read)
	route.RouteGroup.POST("/debt", handler.DebtPayment)
	route.RouteGroup.POST("/debtedit", handler.Edit)          //edit hutang
	route.RouteGroup.GET("/debt", handler.BrowseByCustomer)   //list detail utang
	route.RouteGroup.GET("/debtreport", handler.BrowseByShop) //laporan hutang
}
