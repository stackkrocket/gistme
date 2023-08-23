package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Database connection instance
func DBInstance() *mongo.Client {

	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatalln("Could not load file: ", err)
	}

	log.Println("db conn string active")

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_CONNECTIONSTRING")))
	if err != nil {
		panic("could not establish db client")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		panic("could not connect to database")
	}

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic("could not ping the database")
	}

	log.Println("Pinged db")

	return client

}
