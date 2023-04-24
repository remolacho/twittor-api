package db

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var MongoCnn = connectDB()

// connectDB This function allowed connect to DB in MongoDB
func connectDB() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	uriConn := os.Getenv("DATABASE_URI")
	clientOptions := options.Client().ApplyURI(uriConn)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	return client
}

// CheckConn this function do ping to DB
func CheckConn() int {
	err := MongoCnn.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return 0
	}

	log.Println("the connection is success !!!")

	return 1
}
