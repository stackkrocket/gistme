package helpers

import (
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
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHash(hash, password string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, "Hash and Password do not match! Check your Password again!"
	}
	return true, "Password Match! You may continue Login"
}
