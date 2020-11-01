package bootstrap

import (
	"bukuduit-go/server/bootstrap/routes"
	api "bukuduit-go/server/handlers"
	"net/http"

	"github.com/labstack/echo"
)

func (boot *Bootstrap) RegisterRouters() {
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
		return context.JSON(http.StatusOK, "SIP")
	})

	apiRoute := boot.E.Group("/api/v1")

	//books route
	BooksRoute := apiRoute.Group("/books")
	booksRouteRegistrar := routes.BooksRoute{
		RouteGroup: BooksRoute,
		Handler:    handlerType,
	}
	booksRouteRegistrar.RegisterRoute()

	//author route
	AuthorRoute := apiRoute.Group("/authors")
	authorRouteRegistrar := routes.AuthorRoute{
		RouteGroup: AuthorRoute,
		Handler:    handlerType,
	}
	authorRouteRegistrar.RegisterRoute()

	//category route
	CategoryRoute := apiRoute.Group("/category")
	categoryRouteRegistrar := routes.CategoryRoute{
		RouteGroup: CategoryRoute,
		Handler:    handlerType,
	}
	categoryRouteRegistrar.RegisterRoute()

	//publisher route
	PublisherRoute := apiRoute.Group("/publishers")
	publisherRouteRegistrar := routes.PublisherRoute{
		RouteGroup: PublisherRoute,
		Handler:    handlerType,
	}
	publisherRouteRegistrar.RegisterRoute()

	//barcode route
	BarcodeRoute := apiRoute.Group("/barcode")
	barcodeRouteRegistrar := routes.BarcodeRoute{
		RouteGroup: BarcodeRoute,
		Handler:    handlerType,
	}
	barcodeRouteRegistrar.RegisterRoute()

	//member route
	MemberRoute := apiRoute.Group("/member")
	memberRouteRegistrar := routes.MemberRoute{
		RouteGroup: MemberRoute,
		Handler:    handlerType,
	}
	memberRouteRegistrar.RegisterRoute()

	//bookhascategory route
	BookHasCategoryRoute := apiRoute.Group("/bookCategory")
	bookhascategoryRouteRegistrar := routes.BookHasCategoryRoute{
		RouteGroup: BookHasCategoryRoute,
		Handler:    handlerType,
	}
	bookhascategoryRouteRegistrar.RegisterRoute()

	//book file route
	BookFileRoute := apiRoute.Group("/bookfile")
	bookfileRouteRegistrar := routes.BookFileRoute{
		RouteGroup: BookFileRoute,
		Handler:    handlerType,
	}
	bookfileRouteRegistrar.RegisterRoute()

	//member file route
	MemberFileRoute := apiRoute.Group("/memberfile")
	memberfileRouteRegistrar := routes.MemberFileRoute{
		RouteGroup: MemberFileRoute,
		Handler:    handlerType,
	}
	memberfileRouteRegistrar.RegisterRoute()

	//borrow card route
	BorrowCardRoute := apiRoute.Group("/borrowcard")
	borrowcardRouteRegistrar := routes.BorrowCardRoute{
		RouteGroup: BorrowCardRoute,
		Handler:    handlerType,
	}
	borrowcardRouteRegistrar.RegisterRoute()

}
