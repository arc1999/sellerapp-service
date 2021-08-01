package controller

import (
	"api-scraping-service/model"
	"api-scraping-service/service"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)
//ScrapingController interface
type ScrapingController struct {

}
var scrapingService service.ScrapingService = service.ScrapingServiceImpl{}

//Get ProductDetails
func (sc ScrapingController) Post() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := model.ScrapingRequest{}
		err = c.Bind(&req)
		if err != nil {
			log.Println(err)
			return handleError(err, http.StatusUnprocessableEntity, "Unable To Process Request Body")
		}
		productDetails, err := scrapingService.GetDetails(req)
		if err != nil {
			log.Println(err)
			return handleError(err, http.StatusInternalServerError, "Server error")
		}
		return c.JSON(http.StatusOK, productDetails)
	}
}