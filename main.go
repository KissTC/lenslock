package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// parsing the template and executed it
	// using path
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, "a string")
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

func galleriesByIdHandle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	ctx := r.Context()
	key := ctx.Value("key").(string)

	fmt.Fprint(w, key+id)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handleFunc)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/galleries/{id}", galleriesByIdHandle)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on 3000...")
	http.ListenAndServe(":3000", r)
}
