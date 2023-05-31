package main

import (
	"fmt"
	"net/http"

	"github.com/kisstc/lenslock/controllers"
	"github.com/kisstc/lenslock/models"
	"github.com/kisstc/lenslock/templates"
	"github.com/kisstc/lenslock/views"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// parsing the templates for home page
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))

	// parsing the templates for contact page
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	// parsing the templates for faq page
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(
		templates.FS, "faq.gohtml", "tailwind.gohtml",
	))))

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	userService := models.UserService{
		DB: db,
	}
	// parsing the templates for signup page
	usersC := controllers.Users{
		UserService: &userService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS, "signup.gohtml", "tailwind.gohtml",
	))
	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS, "signin.gohtml", "tailwind.gohtml",
	))
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/users/me", usersC.CurrentUser)

	// r.Use(middleware.Logger)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on 3000...")
	http.ListenAndServe(":3000", r)
}
