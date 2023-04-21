package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/stackkrocket/GistMe/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var users *models.User
var collection *mongo.Client

func CreateUser(name, email, password string) {
	users = &models.User{Name: name, Email: email, Password: password}
	if users.Name == "" || users.Email == "" || users.Password == "" {
		log.Fatalf("%s, %s and %s cannot be empty", users.Name, users.Email, users.Password)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	coll := collection.Database("gistme").Collection("users")
	_, err := coll.InsertOne(ctx, users)
	if err != nil {
		log.Fatalln("Cannot connect to the db!", http.StatusInternalServerError)
	}

	log.Println("Created!", http.StatusCreated)
}
