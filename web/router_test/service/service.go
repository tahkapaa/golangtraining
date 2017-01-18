package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
)

// Person ...
type Person struct {
	Name string
	Age  int
	ID   int
}

// Persons ...
var Persons = []Person{
	Person{"Person 1", 42, 1},
	Person{"Person 2", 14, 2},
	Person{"Person 3", 28, 3},
	Person{"Person 4", 0, 4},
	Person{"Person 5", 54, 5},
}

// PersonToCtx ...
func PersonToCtx(ID string, w http.ResponseWriter, r *http.Request) (context.Context, error) {
	p, err := FindPerson(ID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Println(err)
	}
	return context.WithValue(r.Context(), "person", p), nil
}

// ListPersons ...
func ListPersons(w http.ResponseWriter, r *http.Request) {
	en := json.NewEncoder(w)
	err := en.Encode(Persons)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Println(err)
	}
}

// GetPersonFromCtx ...
func GetPersonFromCtx(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	p, ok := ctx.Value("person").(Person)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		log.Println("Context lost")
		return
	}
	en := json.NewEncoder(w)
	err := en.Encode(p)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Println(err)
	}
}

// FindPerson ...
func FindPerson(ID string) (Person, error) {
	var p Person
	pID, err := strconv.Atoi(ID)
	if err != nil {
		log.Println(err)
		return p, err
	}
	if pID > len(Persons)-1 {
		err = errors.New("Invalid person ID")
		return p, err
	}
	return Persons[pID], nil
}

// GetPerson ...
func GetPerson(p Person, w http.ResponseWriter, r *http.Request) {
	en := json.NewEncoder(w)
	err := en.Encode(p)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Println(err)
	}
}

// AddPerson ...
func AddPerson(w http.ResponseWriter, r *http.Request) {
	de := json.NewDecoder(r.Body)
	var p Person
	err := de.Decode(&p)
	if err != nil {
		log.Println(err)
	}
	p.ID = len(Persons) + 1
	Persons = append(Persons, p)
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Index")
}
