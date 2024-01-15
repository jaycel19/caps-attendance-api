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
	// Attendees routes
	router.Get("/api/v1/attendees", controllers.GetAllAttendees)
	router.Get("/api/v1/attendees/{id}", controllers.GetAttendeeById)
	router.Get("/api/v1/attendees/{course}", controllers.GetAttendeesByCourse)
	router.Post("/api/v1/attendees", controllers.CreateAttendee)
	router.Put("/api/v1/attendees/{id}", controllers.UpdateAttendee)
	router.Delete("/api/v1/attendees/{id}", controllers.DeleteAttendee)

	// Event routes
	router.Get("/api/v1/events", controllers.GetAllEvents)
	router.Get("/api/v1/events/{id}", controllers.GetEventById)
	//router.Get("/api/v1/events/{course}", controllers.GetEventByCourse)
	router.Post("/api/v1/events", controllers.CreateEvent)
	router.Put("/api/v1/events/{id}", controllers.UpdateEvent)
	router.Delete("/api/v1/events/{id}", controllers.DeleteEvent)
	
	// Attendance routes
	router.Get("/api/v1/attendances", controllers.GetAllAttendance)
	router.Get("/api/v1/attendances/{id}", controllers.GetAttendanceById)
	router.Get("/api/v1/attendances/attendee/{id}", controllers.GetAttendanceByAttendeeId)
	router.Get("/api/v1/attendances/event/{id}", controllers.GetAttendanceByEventId)
	router.Get("/api/v1/attendances/{timestart}/{timeend}", controllers.GetAttendanceByRange)
	router.Post("/api/v1/attendances", controllers.CreateAttendance)
	router.Put("/api/v1/attendances/{id}", controllers.UpdateAttendance)
	router.Delete("/api/v1/attendances/{id}", controllers.DeleteAttendance)
	
	// Personnel routes
	router.Get("/api/v1/personnels", controllers.GetAllPersonnel)
	router.Get("/api/v1/personnels/{id}", controllers.GetPersonnelById)
	//router.Get("/api/v1/events/{course}", controllers.GetEventByCourse)
	router.Post("/api/v1/personnels", controllers.CreatePersonnel)
	router.Put("/api/v1/personnels/{id}", controllers.UpdatePersonnel)
	router.Delete("/api/v1/personnels/{id}", controllers.DeletePersonnel)
	

	return router
}
