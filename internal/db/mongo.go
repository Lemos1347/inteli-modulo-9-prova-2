package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBClient struct {
	DBClient *mongo.Client
	DB       *mongo.Database
}

func (s *DBClient) EndDBConnection() {
	if err := s.DBClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func NewDBConnection() *DBClient {
	// uri := os.Getenv("MONGODB_URI")
	uri := "mongodb://root:password@localhost:27017/?retryWrites=true&connectTimeoutMS=10000&authSource=admin&authMechanism=SCRAM-SHA-1&ssl=false"
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return &DBClient{
		DBClient: client,
		DB:       client.Database("messagesstore"),
	}
}
