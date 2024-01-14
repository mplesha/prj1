package server

import (
	"fmt"
	"log"
	"net/http"

	handlers "github.com/RomanTykhyi/students-api/internal/server/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartServer(port int) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/api/v1/students", func(r chi.Router) {
		r.Get("/", handlers.QueryStudents)
		r.Post("/", handlers.CreateStudent)

		// import
		r.Post("/import", handlers.ImportStudents)

		// export
		r.Get("/export", handlers.ExportStudents)

		// subroute
		r.Route("/{studentId}", func(r chi.Router) {
			r.Get("/", handlers.GetStudent)
			r.Put("/", handlers.UpdateStudent)
			r.Delete("/", handlers.DeleteStudent)
		})
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		log.Fatal("Error starting the http server.")
	}
}
