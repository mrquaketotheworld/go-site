package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mrquaketotheworld/go-site/controllers"
	"github.com/mrquaketotheworld/go-site/views"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(views.Must(views.Parse("home"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.Parse("contact"))))
	r.Get("/faq", controllers.StaticHandler(views.Must(views.Parse("faq"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting server on 3000")
	http.ListenAndServe(":3000", r)
}
