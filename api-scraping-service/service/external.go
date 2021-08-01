package service

import (
	"api-scraping-service/model"
	"errors"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"os"
)

type ProductService struct {

}
var restClient = resty.New()

func (s ProductService) SendProductInfo(input model.ProductInfoRequest) (res *model.Product,err error) {
	resp, err := restClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("client", "service").SetResult(&res).
		SetBody(input).Post(os.Getenv("PRODUCT_INFO_URL"))
	if err != nil {
		log.Errorln(err)
	}
	if resp.IsSuccess(){
		return res,errors.New("Request failed")
	}
	return res,nil
}
