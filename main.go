package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/stackkrocket/GistMe/db"
	"github.com/stackkrocket/GistMe/helpers"
	"github.com/stackkrocket/GistMe/routes"
)

func main() {

	helpers.LoadEnv()
	//serves as the main entry point and build of the app
	router := mux.NewRouter()
	db.ConnectDB()

	//serve assets
	fs := http.FileServer(http.Dir("./assets"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	HomePageHandler := helpers.ErrorHandler(routes.HomePage)
	RegisterPageHandler := helpers.ErrorHandler(routes.RegisterPage)
	LoginPageHandler := helpers.ErrorHandler(routes.LoginPage)
	AuthPageHandler := helpers.ErrorHandler(routes.AuthPage)
	RegisterUserHandler := helpers.ErrorHandler(routes.RegisterUser)
	LoginUserHandler := helpers.ErrorHandler(routes.LoginUser)

	router.Handle("/", helpers.CheckHeaderGet(HomePageHandler)).Methods("GET")
	router.Handle("/register", helpers.CheckHeaderGet(RegisterPageHandler)).Methods("GET")
	router.Handle("/login", helpers.CheckHeaderGet(LoginPageHandler)).Methods("GET")
	router.Handle("/auth", helpers.CheckHeaderGet(AuthPageHandler)).Methods("GET")

	router.Handle("/register", helpers.CheckHeaderPost(RegisterUserHandler)).Methods("POST")
	router.Handle("/login", helpers.CheckHeaderPost(LoginUserHandler)).Methods("POST")

	srv := &http.Server{
		Addr:           os.Getenv("PORT"),
		Handler:        router,
		ReadTimeout:    time.Second * 3,
		WriteTimeout:   time.Second * 3,
		MaxHeaderBytes: 1 << 20,
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

	fmt.Println("Server started on PORT", os.Getenv("PORT"))
}
