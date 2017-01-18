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
	s.HandleFunc("/{id}", GetPerson).Methods("GET")

	http.ListenAndServe(":3333", r)
}

// GetPerson ...
func GetPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p, err := service.FindPerson(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Println(err)
	}
	service.GetPerson(p, w, r)
}
