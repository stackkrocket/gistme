package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/stackkrocket/GistMe/helpers"
	"github.com/stackkrocket/GistMe/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/crypto/bcrypt"
)

func HomePage(w http.ResponseWriter, r *http.Request) *helpers.ErrorField {
	path := "./view/index.html"
	responseChan := make(chan []byte)

	go ServeHTMLPage(path, responseChan)
	htmlBytes := <-responseChan
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(htmlBytes)

	return nil
}

func RegisterPage(w http.ResponseWriter, r *http.Request) *helpers.ErrorField {
	path := "./view/register.html"
	responseChan := make(chan []byte)

	go ServeHTMLPage(path, responseChan)
	htmlBytes := <-responseChan
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(htmlBytes)

	return nil
}

func LoginPage(w http.ResponseWriter, r *http.Request) *helpers.ErrorField {
	path := "./view/login.html"
	responseChan := make(chan []byte)

	go ServeHTMLPage(path, responseChan)
	htmlBytes := <-responseChan
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(htmlBytes)

	return nil
}

func AuthPage(w http.ResponseWriter, r *http.Request) *helpers.ErrorField {
	path := "./view/authPage.html"
	responseChan := make(chan []byte)

	go ServeHTMLPage(path, responseChan)
	htmlBytes := <-responseChan
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(htmlBytes)

	return nil
}

func RegisterUser(w http.ResponseWriter, r *http.Request) *helpers.ErrorField {

	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatalln("Could not load file: ", err)
	}

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

	id := primitive.NewObjectID()
	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	password := r.FormValue("password")
	HashedPassword, _ := helpers.PasswordHash(password)
	role := "user"
	verified := false
	user_id := id.Hex()
	access_token, refresh_token, _ := helpers.GenerateAllToken(name, email, user_id)
	created_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	coll := client.Database("gistme").Collection("users")
	var user models.User

	//check for an existing user by email, if found, panic
	//otherwise register user
	filter := bson.D{{Key: "email", Value: email}}
	err = coll.FindOne(ctx, filter).Decode(&user)

	if user.Email == email {
		return &helpers.ErrorField{Error: err, Message: "user already exists. Try login or user another email", Code: 500}
	}

	user = models.User{
		ID:           id,
		Name:         name,
		Email:        email,
		Phone:        phone,
		Password:     HashedPassword,
		Role:         role,
		Verified:     verified,
		User_id:      user_id,
		AccessToken:  access_token,
		RefreshToken: refresh_token,
		CreatedAt:    created_at,
		UpdatedAt:    updated_at,
	}

	if (user.Name != "") && (user.Email != "") && (user.Password != "") {
		result, err := coll.InsertOne(ctx, user)
		if err != nil {
			return &helpers.ErrorField{Error: err, Message: "Internal Server Error: Could not register User", Code: 500}
		}
		fmt.Println(result.InsertedID, user.Password)
		http.Redirect(w, r, "/auth", http.StatusMovedPermanently)
	}
	return nil
}

func LoginUser(w http.ResponseWriter, r *http.Request) *helpers.ErrorField {
	//a struct to unmarshal the returned JSON from the database
	var result *models.User

	err := godotenv.Load("app.env")
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

	access_token, refresh_token, _ := helpers.GenerateAllToken(result.Name, result.Email, result.User_id)

	if bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		return &helpers.ErrorField{Error: err, Message: "", Code: 500}
	}

	helpers.UpdateAllToken(access_token, refresh_token, result.User_id)

	log.Println("Authentication successful")
	http.Redirect(w, r, "/auth", http.StatusMovedPermanently)
	return nil
}
