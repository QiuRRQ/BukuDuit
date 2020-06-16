package middleware

import (
	"fmt"
	"github.com/labstack/echo"
	"bukuduit-go/helpers/logrus"
	"net/http"
)

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logrus.MakeLogEntry(c).Info("incoming request")
		return next(c)
	}
}

func errorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	c.Response()
	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	logrus.MakeLogEntry(c).Error(report.Message)
	c.HTML(report.Code, report.Message.(string))
}
