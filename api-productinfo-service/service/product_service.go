package service

import (
	"api-productinfo-service/dao"
	"api-productinfo-service/model"
)

//ProductService interface
type ProductService interface {
	Save(input model.ProductRequest) (*model.Product, error)
}

//ProductService interface
type ProductServiceImpl struct {
}

var productDao dao.ProductDao = dao.ProductDaoImpl{}

//Get & Process Product Details
func (i ProductServiceImpl) Save(input model.ProductRequest) (*model.Product, error) {
	var req model.Product
	req.Url = input.Url
	req.ProductInfo = input.ProductInfo
	return productDao.Save(req)
}
