package routes

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/stackkrocket/GistMe/helpers"
	"github.com/stackkrocket/GistMe/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var users *models.User

// var collection *mongo.Client
var tmpl, err = template.ParseGlob("./view/*.html")

func HomePage(w http.ResponseWriter, r *http.Request) *helpers.ErrorField {
	w.Header().Set("Content-Type", "text/html")
	err = tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		return &helpers.ErrorField{Error: err, Message: "error parsing template", Code: 404}
	}
	return nil
}

func RegisterPage(w http.ResponseWriter, r *http.Request) *helpers.ErrorField {
	w.Header().Set("Content-Type", "text/html")
	err = tmpl.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		return &helpers.ErrorField{Error: err, Message: "error parsing template", Code: 404}
	}
	return nil
}

func LoginPage(w http.ResponseWriter, r *http.Request) *helpers.ErrorField {
	w.Header().Set("Content-Type", "text/html")
	err = tmpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		return &helpers.ErrorField{Error: err, Message: "error parsing template", Code: 404}
	}
	return nil
}

func AuthPage(w http.ResponseWriter, r *http.Request) *helpers.ErrorField {
	w.Header().Set("Content-Type", "text/html")
	err = tmpl.ExecuteTemplate(w, "authPage.html", nil)
	if err != nil {
		return &helpers.ErrorField{Error: err, Message: "error parsing template", Code: 404}
	}
	return nil
}

func RegisterUser(w http.ResponseWriter, r *http.Request) *helpers.ErrorField {
	helpers.LoadEnv()
	//Establish atlas cluster connection here
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_CONNECTIONSTRING")))
	if err != nil {
		return &helpers.ErrorField{Error: err, Message: "server error: could not establish a connection with the server", Code: 500}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		return &helpers.ErrorField{Error: err, Message: "Could not establish connection with database", Code: 500}
	}

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return &helpers.ErrorField{Error: err, Message: "Database took too long to respond", Code: 500}
	}

	r.ParseForm()

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	//hash the password before being saved to db
	HashedPassword, _ := helpers.PasswordHash(password)
	//fmt.Println(HashedPassword)

	coll := client.Database(os.Getenv("DBNAME")).Collection("users")
	users = &models.User{Name: name, Email: email, Password: HashedPassword}

	if (users.Name != "") && (users.Email != "") && (users.Password != "") {
		result, err := coll.InsertOne(ctx, users)
		if err != nil {
			return &helpers.ErrorField{Error: err, Message: "Internal Server Error: Could not register User", Code: 500}
		}
		fmt.Println(result.InsertedID)
		http.Redirect(w, r, "/auth", http.StatusMovedPermanently)
	}
	return nil
}

func LoginUser(w http.ResponseWriter, r *http.Request) *helpers.ErrorField {
	helpers.LoadEnv()
	//Establish atlas cluster connection here
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_CONNECTIONSTRING")))
	if err != nil {
		return &helpers.ErrorField{Error: err, Message: "server error: could not establish a connection with the server", Code: 500}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		return &helpers.ErrorField{Error: err, Message: "Could not establish connection with database", Code: 500}
	}

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return &helpers.ErrorField{Error: err, Message: "Database took too long to respond", Code: 500}
	}

	r.ParseForm()

	email := r.FormValue("email")
	password := r.FormValue("password")

	//fmt.Println(email)
	coll := client.Database(os.Getenv("DBNAME")).Collection("users")
	filter := bson.D{{Key: "email", Value: email}}
	var result *models.User

	err = coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &helpers.ErrorField{Error: err, Message: "User does not exist", Code: 500}
		}
		log.Println("Internal server error")
	}
	//fmt.Println(result.Password)
	match, _ := helpers.CompareHash(result.Password, password)

	if match {
		http.Redirect(w, r, "/auth", http.StatusAccepted)
		return nil
	}
	return &helpers.ErrorField{Error: err, Message: "Password Incorrect", Code: 500}
}
