package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/tahkapaa/golangtraining/web/router_test/service"
)

func main() {
	h := chi.NewRouter()
	h.Use(middleware.Logger)
	h.HandleFunc("/", service.Index)
	h.Route("/persons", func(r chi.Router) {
		r.Get("/", service.ListPersons)
		r.Post("/persons", service.AddPerson)
		r.Route("/:personID", func(sr chi.Router) {
			sr.Use(personCtx)
			sr.Get("/", service.GetPerson)
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
		log.Printf("PersonID: %d, personcount: %d", personID, len(service.Persons))
		if personID > len(service.Persons)-1 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			log.Println("Invalid person ID")
			return
		}
		p := service.Persons[personID]
		ctx := context.WithValue(r.Context(), "person", p)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
