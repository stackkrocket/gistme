package helpers

import (
	"encoding/base64"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// define error custome fields (properties) to log to the debugger
type ErrorField struct {
	Error   error
	Message string
	Code    int
}

// custom error Handler of type http.Handler
type ErrorHandler func(w http.ResponseWriter, r *http.Request) *ErrorField

// the http package doesn’t understand functions that return error.
// implementing http.Handler ServerHTTP interface to reduce the verbose err != nil syntax
func (er ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := er(w, r); e != nil { //e is of type *ErrorField not type os.Error
		log.Printf("%v", e.Error)
		http.Error(w, e.Message, e.Code)
	}
}

// Load environment variables
func LoadEnv() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Could not load file: ", err)
	}

	apiKey := os.Getenv("DB_CONNECTIONSTRING")
	return apiKey
}

func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
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

/*func CompareHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Fatalln(err)
		return false
	}
	log.Println("Success! You may proceed!...")
	return true
}*/
