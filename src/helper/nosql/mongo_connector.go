package nosql

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mesutpiskin.com/gostock/helper"
	"mesutpiskin.com/gostock/model"
)

//MongoClient db accesor
var MongoClient *mongo.Client

// ConnectToMongoDB get db connection
func ConnectToMongoDB() *mongo.Client {
	if MongoClient != nil {
		return MongoClient
	}
	clientOptions := options.Client().ApplyURI(getConnectionString())
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Mongo database connection successful.")
	MongoClient = client
	return client
}

func getConnectionString() string {
	var configurationModel model.ConfigurationModel
	configurationModel = helper.GetConfiguration()
	if configurationModel.Mongo.UserName == "" ||
		configurationModel.Mongo.Password == "" ||
		configurationModel.Mongo.UserName == "" {
		log.Fatal("MongoDB configuration error.")
		return ""
	}
	// Initialize mongo connection
	var connection string = "mongodb+srv://%s:%s@%s?retryWrites=true&w=majority"
	dbuser := configurationModel.Mongo.UserName
	dbpass := configurationModel.Mongo.Password
	host := configurationModel.Mongo.HostURL

	mongoDbConnectionString := fmt.Sprintf(connection, dbuser, dbpass, host)
	return mongoDbConnectionString
}
