package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tahkapaa/golangtraining/web/router_test/service"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", service.Index)
	r.HandleFunc("/persons", service.ListPersons).Methods("GET")
	r.HandleFunc("/persons", service.AddPerson).Methods("POST")

	s := r.PathPrefix("/persons").Subrouter()
	s.HandleFunc("/{id}", getPerson).Methods("GET")
	fs := s.PathPrefix("/{id}").Subrouter()
	fs.HandleFunc("/friends", listFriends).Methods("GET")

	http.ListenAndServe(":3333", r)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p, err := service.FindPerson(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Println(err)
		return
	}
	service.PersonToJSON(p, w, r)
}

func listFriends(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p, err := service.FindPerson(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Println(err)
		return
	}
	service.ListFriends(p, w, r)
}
