package dao

import (
	"api-productinfo-service/db"
	"api-productinfo-service/model"
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProductDao interface {
	Save(input model.Product) (*model.Product, error)
}
type ProductDaoImpl struct {
}

//GetAll -
func (lr ProductDaoImpl) Save(input model.Product) (*model.Product, error) {
	var res model.Product
	product, err := getByUrl(input.Url)
	if  err==nil{
		collection := db.GetCollection().Products
		product.UpdatedAt = time.Now()
		product.ProductInfo = input.ProductInfo
		err := collection.FindOneAndReplace(context.TODO(), bson.M{"url": product.Url}, product).Decode(&res)
		if err != nil {
			log.Printf("error in saving Plan")
			return nil, err
		}
		return product, nil
	}else {
		collection := db.GetCollection().Products
		t := time.Now()
		input.CreatedAt = t
		input.UpdatedAt = t
		result, err := collection.InsertOne(context.TODO(), input)
		if err != nil {
			log.Errorln("Error Occured Insertone")
			return nil, err
		}
		insertedID := result.InsertedID.(primitive.ObjectID)
		saved, err := getById(insertedID)
		if err != nil {
			log.Errorln("Error Occured while retrieve")
			return nil, err
		}
		return saved, nil
	}
}
func getByUrl(url string) (*model.Product, error) {
	var p model.Product
	collection := db.GetCollection().Products
	filter := bson.M{"url": url}
	err := collection.FindOne(context.TODO(), filter).Decode(&p)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	return &p, nil
}
func getById(id primitive.ObjectID) (*model.Product, error) {
	var p model.Product
	collection := db.GetCollection().Products
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&p)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	return &p, nil
}
