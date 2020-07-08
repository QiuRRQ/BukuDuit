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

//transaction route
func (route TransactionRoute) RegisterRoute() {
	handler := api.TransactionHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	route.RouteGroup.Use(jwtMiddleware.JWTWithConfig)
	route.RouteGroup.GET("/list", handler.TransactionList)
	route.RouteGroup.GET("/transactionreport", handler.TransactionReport)
	route.RouteGroup.GET("/transaction/file", handler.TransactionReportExportFile)
	route.RouteGroup.DELETE("/:id", handler.DeleteDebt)
	route.RouteGroup.DELETE("/deletetrans/:id", handler.DeleteTrans)
	route.RouteGroup.POST("/transaction", handler.AddTransaction)     //done
	route.RouteGroup.POST("/transactionedit", handler.EditTransction) //change to edit
	route.RouteGroup.GET("/details/:id", handler.Read)
	route.RouteGroup.POST("/debt", handler.DebtPayment)
	route.RouteGroup.POST("/debtedit", handler.Edit)                       //edit hutang
	route.RouteGroup.GET("/debt", handler.BrowseByCustomer)                //list detail utang
	route.RouteGroup.GET("/debt/file", handler.DebtDetailExportFile)       // untuk utang detail file
	route.RouteGroup.GET("/debtreport", handler.BrowseByShop)              //laporan hutang
	route.RouteGroup.GET("/debtreport/file", handler.DebtReportExportFile) //laporan utang file
}
