package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jaycel19/capstone-api/controllers"
)

func Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	// Students routes
	router.Get("/api/v1/students", controllers.GetAllStudents)
	router.Get("/api/v1/students/{id}", controllers.GetStudentById)
	router.Get("/api/v1/students/{course}", controllers.GetStudentsByCourse)
	router.Post("/api/v1/students", controllers.CreateStudent)
	router.Put("/api/v1/students/{id}", controllers.UpdateStudent)
	router.Delete("/api/v1/students/{id}", controllers.DeleteStudent)

	// Event routes
	router.Get("/api/v1/events", controllers.GetAllEvents)
	router.Get("/api/v1/events/{id}", controllers.GetEventById)
	//router.Get("/api/v1/events/{course}", controllers.GetEventByCourse)
	router.Post("/api/v1/events", controllers.CreateEvent)
	router.Put("/api/v1/events/{id}", controllers.UpdateEvent)
	router.Delete("/api/v1/events/{id}", controllers.DeleteEvent)

	return router
}
