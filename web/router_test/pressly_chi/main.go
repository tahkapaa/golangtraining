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
		r.Post("/persons", service.AddPerson)
		r.Route("/:personID", func(sr chi.Router) {
			sr.Use(PersonCtx)
			sr.Get("/", service.GetPersonFromCtx)
		})
	})
	http.ListenAndServe(":3301", h)
}

// PersonCtx ...
func PersonCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		personID := chi.URLParam(r, "personID")
		ctx, err := service.PersonToCtx(personID, w, r)
		if err != nil {
			log.Println(err)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
