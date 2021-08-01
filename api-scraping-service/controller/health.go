package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

//HealthController interface
type HealthController struct{}

type ApplicationError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

//Status - return status
func (h HealthController) Status() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.String(http.StatusOK, "Working!")
	}
}

//HandleError - handle error
func handleError(err error, code int, message string) error {
	log.Errorln(err)
	appError := ApplicationError{Code: code, Message: message}
	log.Errorln(appError)
	return echo.NewHTTPError(code, appError)
}
