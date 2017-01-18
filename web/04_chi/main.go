package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

type person struct {
	Name string
	Age  int
	ID   int
}

var persons = []person{
	person{"Person 1", 42, 1},
	person{"Person 2", 14, 2},
	person{"Person 3", 28, 3},
	person{"Person 4", 0, 4},
	person{"Person 5", 54, 5},
}

func main() {
	h := chi.NewRouter()
	h.Use(middleware.Logger)
	h.HandleFunc("/", index)
	h.Get("/foo", foo)
	h.HandleFunc("/bar", bar)
	h.Route("/persons", func(r chi.Router) {
		r.Get("/", listPersons)
		r.Post("/persons", addPerson)
		r.Route("/:personID", func(sr chi.Router) {
			sr.Use(personCtx)
			sr.Get("/", getPerson)
		})
	})
	http.ListenAndServe(":3301", h)
}

func personCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		personID, err := strconv.Atoi(chi.URLParam(r, "personID"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			log.Println(err)
			return
		}
		log.Printf("PersonID: %d, personcount: %d", personID, len(persons))
		if personID > len(persons)-1 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			log.Println(err)
			return
		}
		p := persons[personID]
		ctx := context.WithValue(r.Context(), "person", p)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	p, ok := ctx.Value("person").(person)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		log.Println("Context lost")
	}
	en := json.NewEncoder(w)
	err := en.Encode(p)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Println(err)
	}
}

func listPersons(w http.ResponseWriter, r *http.Request) {
	en := json.NewEncoder(w)
	err := en.Encode(persons)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Println(err)
	}
}

func addPerson(w http.ResponseWriter, r *http.Request) {
	de := json.NewDecoder(r.Body)
	var p person
	err := de.Decode(&p)
	if err != nil {
		log.Println(err)
	}
	p.ID = len(persons) + 1
	persons = append(persons, p)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Index")
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo")
}
