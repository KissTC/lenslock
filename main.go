package main

import (
	"fmt"
	"net/http"

	"github.com/kisstc/lenslock/templates"

	"github.com/kisstc/lenslock/controllers"

	"github.com/kisstc/lenslock/views"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// parsing the templates for home page
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml"))))

	// parsing the templates for contact page
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))

	// parsing the templates for faq page
	r.Get("/faq", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))

	// r.Use(middleware.Logger)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on 3000...")
	http.ListenAndServe(":3000", r)
}
