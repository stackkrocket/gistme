package helpers

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/crypto/bcrypt"
)

// define error custome fields (properties) to log to the debugger
type ErrorField struct {
	Error   error
	Message string
	Code    int
}

type SignedDetails struct {
	Email string
	Name  string
	Uid   string
	jwt.StandardClaims
}

// custom error Handler of type http.Handler
type ErrorHandler func(w http.ResponseWriter, r *http.Request) *ErrorField

// var uc *mongo.Collection = db.OpenCollection(db.Client, "user")
var SECRET_KEY string = os.Getenv("SECRET_KEY")

// the http package doesn’t understand functions that return error.
// implementing http.Handler ServerHTTP interface to reduce the verbose err != nil syntax
func (er ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := er(w, r); e != nil { //e is of type *ErrorField not type os.Error
		log.Printf("%v", e.Error)
		http.Error(w, e.Message, e.Code)
	}
}

// Load environment variables
func LoadEnv() {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatalln("Could not load file: ", err)
	}
	fmt.Println("Loaded env file")
}

func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "Sorry Password could not be hashed...", err
	}

	/*
		it appears mongodb requires a standard encoding to actually store the hash in the correct format
		without the encoding/base64 package, when you try to login, and the function
		attempts to compare supposed stored to the plain text
		equivalent, it returns a weird 'HashedSecret too small to be a bcrypt password".
		At this moment, this seems to work. Althouh,
		I am putting in more research to ensure there was not another reason for the error
	*/
	encodedPassword := base64.StdEncoding.EncodeToString(bytes)
	return string(encodedPassword), nil
}

func GenerateAllToken(email, name, uid string) (signedAccessToken, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email: email,
		Name:  name,
		Uid:   uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * 3600).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * 3600).Unix(),
		},
	}

	access_token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refresh_token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
	}

	return access_token, refresh_token, nil
}

// validate generated tokens
func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "the token is invalid"
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token has expired"
		return
	}

	return claims, "Token Validated...!!!"
}

//upon successful signin, update user tokens

func UpdateAllToken(signedAccessToken, signedRefreshToken, userId string) {
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

	var updateObj primitive.D

	coll := client.Database("gistme").Collection("users")

	updateObj = append(updateObj, bson.E{Key: "access_token", Value: signedAccessToken})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: signedRefreshToken})

	UpdatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: UpdatedAt})

	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	result, err := coll.UpdateOne(
		ctx,
		filter,
		bson.D{{Key: "$set", Value: updateObj}},
		&opt,
	)

	defer cancel()

	if err != nil {
		log.Panic(err)
		return
	}

	fmt.Println("Updated new user", result)
}
