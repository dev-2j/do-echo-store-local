package connx

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var Client *mongo.Client
type DtoMg struct {
	Client *mongo.Client
	Cnx    *mongo.Collection
}

func Mg() *DtoMg {
	return &DtoMg{}
}

func (mg *DtoMg) ConnextMongo() error {
	// connext mongo
	dsn := `mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.3.0`
	if dsn == "" {
		log.Fatal("MONGO_DB environment variable is missing")
	}

	var err error
	mg.Client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Check connection
	err = mg.Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
	}

	return nil
}
