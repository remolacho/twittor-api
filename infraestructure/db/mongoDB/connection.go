package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
)

var (
	data *Mongo
	once sync.Once
)

type Mongo struct {
	Client *mongo.Client
}

func CurrentSession() *Mongo {
	once.Do(initDB)
	return data
}

func (m *Mongo) Check() int {
	err := m.Client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return 0
	}

	return 1
}

func (m *Mongo) DataBase() *mongo.Database {
	return data.Client.Database(os.Getenv("DATABASE_NAME"))
}

func initDB() {
	client := getConnection()
	data = &Mongo{Client: client}
}

func getConnection() *mongo.Client {
	uriConn := os.Getenv("DATABASE_URI")
	clientOptions := options.Client().ApplyURI(uriConn)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("the connection to mongoDB is success !!!")

	return client
}
