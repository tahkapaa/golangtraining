package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/tahkapaa/golangtraining/web/exercises/08_mongodb/08_handson/starting-code/models"
)

type UserController struct {
	users map[string]models.User
}

func NewUserController() *UserController {
	uc := UserController{}
	uc.users = make(map[string]models.User)
	uc.loadData()

	return &uc
}

func (uc *UserController) loadData() {
	file, err := os.Open("users")
	if err != nil {
		file, err = os.Create("users")
		if err != nil {
			log.Fatalln(err)
		}
		file.Close()
		fmt.Println(err)
		return
	}
	dec := json.NewDecoder(file)
	for dec.More() {
		var u models.User
		if err := dec.Decode(&u); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(u)
		uc.users[u.Id] = u
	}
	file.Close()
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")
	fmt.Println(id)

	// composite literal
	u, ok := uc.users[id]

	// Fetch user
	if !ok {
		w.WriteHeader(404)
		return
	}
	fmt.Println(u)

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create bson ID
	id, err := uuid.NewV4()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	u.Id = id.String()
	uc.users[u.Id] = u
	uj, _ := json.Marshal(u)

	file, err := os.OpenFile("users", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.FileMode(0600))
	if err != nil {
		log.Fatalln(err)
	}
	if err := json.NewEncoder(file).Encode(u); err != nil {
		log.Fatalln(err)
	}
	file.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	delete(uc.users, id)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")
}
