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
	"golang.org/x/crypto/bcrypt"
)

var users models.User

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

	HashedPassword, err := helpers.PasswordHash(password)
	if err != nil {
		panic("Password could not be hashed...")
	}

	coll := client.Database(os.Getenv("DBNAME")).Collection("users")
	users = models.User{Name: name, Email: email, Password: HashedPassword}

	if (users.Name != "") && (users.Email != "") && (users.Password != "") {
		result, err := coll.InsertOne(ctx, users)
		if err != nil {
			return &helpers.ErrorField{Error: err, Message: "Internal Server Error: Could not register User", Code: 500}
		}
		fmt.Println(result.InsertedID, users.Password)
		http.Redirect(w, r, "/auth", http.StatusMovedPermanently)
	}
	return nil
}

func LoginUser(w http.ResponseWriter, r *http.Request) *helpers.ErrorField {
	//a struct to unmarshal the returned JSON from the database
	var result *models.User

	helpers.LoadEnv()
	//Establish atlas cluster connection here
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_CONNECTIONSTRING")))
	if err != nil {
		log.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		log.Println(err)
	}

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err)
	}

	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	//retrieve a document providing the email
	coll := client.Database("gistme").Collection("users")
	filter := bson.D{{Key: "email", Value: email}}

	err = coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Fatalln("Status 500: Could not find any matching documents")
		}
	}
	fmt.Println(result.Email, result.Password)

	/*if result.Email == email && err == nil {
		log.Println("Authentication Succesful!")
		http.Redirect(w, r, "/auth", http.StatusMovedPermanently)
	}

	log.Println("Email or password do not match")*/

	if bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		return &helpers.ErrorField{Error: err, Message: "", Code: 500}
	}

	log.Println("Authentication successful")
	http.Redirect(w, r, "/auth", http.StatusMovedPermanently)
	return nil
}
