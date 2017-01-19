package main

import (
	"log"
	"net/http"

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
		r.Post("/", service.AddPerson)
		r.Route("/:personID", func(sr chi.Router) {
			sr.Use(PersonCtx)
			sr.Get("/", service.GetPersonFromCtx)
			sr.Get("/friends", listFriends)
		})
	})
	http.ListenAndServe(":3333", h)
}

// PersonCtx ...
func PersonCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		personID := chi.URLParam(r, "personID")
		ctx, err := service.PersonToCtx(personID, w, r)
		if err != nil {
			log.Println(err)
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func listFriends(w http.ResponseWriter, r *http.Request) {
	p, err := service.PersonFromCtx(r)
	if err != nil {
		log.Println(err)
		return
	}
	service.ListFriends(p, w, r)
}
