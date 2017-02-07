package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tahkapaa/golangtraining/web/exercises/08_mongodb/08_handson/starting-code/controllers"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	log.Fatalln(http.ListenAndServe("localhost:8080", r))
}
