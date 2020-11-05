package nosql

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToMongoDB get db connection
func ConnectToMongoDB() *mongo.Client {
	// Initialize mongo connection
	var connection string = "mongodb+srv://%s:%s@%s?retryWrites=true&w=majority"
	dbuser := os.Getenv("MONGO_USER")
	dbpass := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")

	mongoDbConnectionString := fmt.Sprintf(connection, dbuser, dbpass, host)

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoDbConnectionString)

	// Connect to MongoDB
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

	return client
}
