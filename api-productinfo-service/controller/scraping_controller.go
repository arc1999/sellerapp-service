package controller

import (
	"api-productinfo-service/model"
	"api-productinfo-service/service"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

//ScrapingController interface
type ProductController struct {
}

var productService service.ProductService = service.ProductServiceImpl{}

//Get ProductDetails
func (sc ProductController) Post() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := model.ProductRequest{}
		err = c.Bind(&req)
		if err != nil {
			log.Println(err)
			return handleError(err, http.StatusUnprocessableEntity, "Unable To Process Request Body")
		}
		productDetails, err := productService.Save(req)
		if err != nil {
			log.Println(err)
			return handleError(err, http.StatusInternalServerError, "Server error")
		}
		return c.JSON(http.StatusOK, productDetails)
	}
}
