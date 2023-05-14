package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/kisstc/lenslock/controllers"

	"github.com/kisstc/lenslock/views"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// parsing the templates for home page
	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	// parsing the templates for contact page
	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	// parsing the templates for faq page
	tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	// r.Use(middleware.Logger)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on 3000...")
	http.ListenAndServe(":3000", r)
}
