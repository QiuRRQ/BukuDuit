package bootstrap

import (
	api "bukuduit-go/server/handlers"
	"github.com/labstack/echo"
	"net/http"
)

func(boot *Bootstrap) RegisterRouters(){
	_ = api.Handler{
		E:               boot.E,
		Db:              boot.Db,
		UseCaseContract: &boot.UseCaseContract,
		Jwe:             boot.Jwe,
		Validate:        boot.Validator,
		Translator:      boot.Translator,
		JwtConfig:       boot.JwtConfig,
	}

	boot.E.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "SIP")
	})
<<<<<<< HEAD
}
=======

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

	//transaction route
	transactionRoute := apiRoute.Group("/transactions")
	transactionRouteRegistrar := routes.TransactionRoute{
		RouteGroup: transactionRoute,
		Handler:    handlerType,
	}
	transactionRouteRegistrar.RegisterRoute()

	userRoute := apiRoute.Group("/user")
	userRouteRegister := routes.UserRoute{
		RouteGroup: userRoute,
		Handler:    handlerType,
	}
	userRouteRegister.RegisterRoute()

	payAccRoute := apiRoute.Group("/payacc")
	payAccRouteRegister := routes.PaymentAccountRoute{
		RouteGroup: payAccRoute,
		Handler:    handlerType,
	}
	payAccRouteRegister.RegisterRoute()
}
>>>>>>> stage
