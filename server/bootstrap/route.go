package bootstrap

import (
	"bukuduit-go/server/bootstrap/routes"
	api "bukuduit-go/server/handlers"
	"github.com/labstack/echo"
	"net/http"
)

func(boot *Bootstrap) RegisterRouters(){
	handlerType := api.Handler{
		E:               boot.E,
		Db:              boot.Db,
		UseCaseContract: &boot.UseCaseContract,
		Jwe:             boot.Jwe,
		Validate:        boot.Validator,
		Translator:      boot.Translator,
		JwtConfig:       boot.JwtConfig,
	}

	boot.E.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "Work")
	})

	apiRoute := boot.E.Group("/api/v1")

	//otp route
	otpRoute := apiRoute.Group("/otp")
	otpRouteRegistrar := routes.OtpRoutes{
		RouteGroup: otpRoute,
		Handler:    handlerType,
	}
	otpRouteRegistrar.RegisterRoute()

	//authentication route
	authenticationRoute := apiRoute.Group("/auth")
	authenticationRegistrar := routes.AuthenticationRoute{
		RouteGroup: authenticationRoute,
		Handler:    handlerType,
	}
	authenticationRegistrar.RegisterRoute()

	//businesscard route
	shopRoute := apiRoute.Group("/shop")
	shopRouteRegistrar := routes.ShopRoute{
		RouteGroup: shopRoute,
		Handler:    handlerType,
	}
	shopRouteRegistrar.RegisterRoute()

	//usercustomer route
	userCustomerRoute := apiRoute.Group("/user-customer")
	userCustomerRouteRegistrar := routes.UserCustomerRoute{
		RouteGroup: userCustomerRoute,
		Handler:    handlerType,
	}
	userCustomerRouteRegistrar.RegisterRoute()
}