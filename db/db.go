package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/stackkrocket/GistMe/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	Name     string
	Email    string
	Password string
}

func ConnectDB() {
	helpers.LoadEnv()
	//Establish atlas cluster connection here
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_CONNECTIONSTRING")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Could not find database", err)
	}
	log.Println("Connected to database successfully...!")

	//insert a document
	/*coll := client.Database("gistme").Collection("users")
	users := []interface{}{
		User{Name: "Yusuf", Email: "rabiu@gmail.com", Password: "123456"},
		User{Name: "Bahirah", Email: "bahirah@gmail.com", Password: "1525i6"},
		User{Name: "Rabiu", Email: "yusuf@gmail.com", Password: "r8fieaklsf"},
	}
	result, err := coll.InsertMany(ctx, users)
	if err != nil {
		log.Fatal("Cannot add document!", err)
	}
	for _, id := range result.InsertedIDs {
		fmt.Println("Inserted documents with IDs: ", id)
	}*/

}
