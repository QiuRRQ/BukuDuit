package routes

import (
	api "bukuduit-go/server/handlers"

	"github.com/labstack/echo"
)

type BorrowCardRoute struct {
	RouteGroup *echo.Group
	Handler    api.Handler
}

func (route BorrowCardRoute) RegisterRoute() {
	handler := api.BorrowCardHandler{Handler: route.Handler}
	// jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	// route.RouteGroup.Use(jwtMiddleware.JWTWithConfig)
	route.RouteGroup.GET("/data", handler.Read)
	route.RouteGroup.GET("/:id", handler.ReadByID)
	route.RouteGroup.POST("", handler.Add)
	route.RouteGroup.POST("/peminjaman", handler.Pinjaman)
	route.RouteGroup.GET("/laporanpinjam", handler.LaporanPinjam)
	route.RouteGroup.POST("/pengembalian", handler.Kembalikan)
	route.RouteGroup.GET("/laporanpengembalian", handler.LaporanPengembalian)
	route.RouteGroup.PUT("/:id", handler.Edit)
	route.RouteGroup.DELETE("/:id", handler.Delete)
}
