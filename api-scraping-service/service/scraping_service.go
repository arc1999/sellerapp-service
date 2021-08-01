package service

import (
	"api-scraping-service/model"
	"github.com/gocolly/colly"
	"log"
	"strconv"
	"strings"
)
//ScrapingService interface
type ScrapingService interface {
	GetDetails(input model.ScrapingRequest) (p model.ProductDetailsRequest,err error)
}
//ScrapingService interface
type ScrapingServiceImpl struct {
	
}
var externalService ProductService
//Get & Process Product Details
func (i ScrapingServiceImpl) GetDetails(input model.ScrapingRequest) (model.ProductDetailsRequest, error) {
	product:=GetAmazonProductDetails(input.Url)
	externalService.SendProductInfo(model.ProductInfoRequest{
		Url:         input.Url,
		ProductInfo: &product,
	})
	return product,nil
}
//Scrap from Amazon's Url
func GetAmazonProductDetails(url string) (p model.ProductDetailsRequest) {
	c := colly.NewCollector(
		colly.AllowedDomains("amazon.in", "www.amazon.in"),
	)

	c.OnRequest(func(r *colly.Request) {
	})
	c.OnError(func(response *colly.Response, err error) {
		log.Println(err)
	})

	// Getting product title using the 'id' attribute of 'span' tag
	c.OnHTML("span[id=productTitle]", func(e *colly.HTMLElement) {
		p.Name = strings.Trim(e.Text, "\n")
	})

	// Getting the review count similarly
	c.OnHTML("span[id=acrCustomerReviewText]", func(e *colly.HTMLElement) {
		cnt, _ := strconv.ParseInt(strings.Replace(strings.SplitN(e.Text, " ", 2)[0], ",", "", -1), 10, 64)
		p.Reviews = int(cnt)
	})

	// Getting product description and storing it as array of strings
	c.OnHTML("div[id=feature-bullets]", func(e *colly.HTMLElement) {
		e.ForEach("ul", func(_ int, ul *colly.HTMLElement) {
			log.Println(ul.Text)
					p.Description =  ul.Text
			})
		})
	// Getting product price
	c.OnHTML("span[id=priceblock_ourprice]", func(e *colly.HTMLElement) {
		p.Price = e.Text
		//if p.Price==""{
		//	c.OnHTML("span[id=priceblock_dealprice]", func(e *colly.HTMLElement) {
		//		p.Price = e.Text
		//	})
		//}
	})

	// Getting image url
	c.OnHTML("div[class=imgTagWrapper]", func(e *colly.HTMLElement) {
		e.ForEach("img[src]", func(_ int, i *colly.HTMLElement) {
			p.ImagesURL = i.Attr("src")
		})
	})
	c.Visit(url)
	return p
}
