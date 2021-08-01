package db

import (
	"context"
	"fmt"
	"os"
	"reflect"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var db *mongo.Database
var collection *mongo.Collection

// InitDb initialize DB...
func InitDb() {
	host := os.Getenv("DATABASE_URI")
	rb := bson.NewRegistryBuilder()
	rb.RegisterTypeMapEntry(bsontype.EmbeddedDocument, reflect.TypeOf(bson.M{}))

	clientOptions := options.Client().ApplyURI(host).SetRegistry(rb.Build())

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Panicln(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {

		log.Panicln(err)
	}
	db = client.Database(os.Getenv("MONGO_DB_NAME"))
	fmt.Println("Connected to MongoDB!")
}

// GetDB returns DB instance
func GetDB() *mongo.Database {
	return db
}

// GetCollection returns Collection instance
func GetCollection() MongoCollection {
	return MongoCollection{
		Products: GetDB().Collection(os.Getenv("COLLECTION_NAME")),
	}
}

type MongoCollection struct {
	Products *mongo.Collection
}
