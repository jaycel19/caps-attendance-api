package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jaycel19/capstone-api/controllers"
	"github.com/jaycel19/capstone-api/middlewares"
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

	router.Post("/api/v1/admin/login", controllers.LoginAdmin)
	router.Post("/api/v1/personnels/login", controllers.LoginPersonnel)
	router.Post("/api/v1/personnels", controllers.CreatePersonnel)
	router.Post("/api/v1/admin", controllers.CreateAdmin)

	router.Group(func(auth chi.Router){
		auth.Use(middlewares.RequireAuth)

		// Attendees routes
		auth.Get("/api/v1/attendees", controllers.GetAllAttendees)
		auth.Get("/api/v1/attendees/{id}", controllers.GetAttendeeById)
		//auth.Get("/api/v1/attendees/{course}", controllers.GetAttendeesByCourse)
		auth.Post("/api/v1/attendees", controllers.CreateAttendee)
		auth.Put("/api/v1/attendees/{id}", controllers.UpdateAttendee)
		auth.Delete("/api/v1/attendees/{id}", controllers.DeleteAttendee)

		// Event routes
		auth.Get("/api/v1/events", controllers.GetAllEvents)
		auth.Get("/api/v1/events/{id}", controllers.GetEventById)
		//auth.Get("/api/v1/events/{course}", controllers.GetEventByCourse)
		auth.Post("/api/v1/events", controllers.CreateEvent)
		auth.Put("/api/v1/events/{id}", controllers.UpdateEvent)
		auth.Delete("/api/v1/events/{id}", controllers.DeleteEvent)
		
		// Attendance routes
		auth.Get("/api/v1/attendances", controllers.GetAllAttendance)
		auth.Get("/api/v1/attendances/{id}", controllers.GetAttendanceById)
		auth.Get("/api/v1/attendances/attendee/{id}", controllers.GetAttendanceByAttendeeId)
		auth.Get("/api/v1/attendances/event/{id}", controllers.GetAttendanceByEventId)
		auth.Get("/api/v1/attendances/{timestart}/{timeend}", controllers.GetAttendanceByRange)
		auth.Post("/api/v1/attendances", controllers.CreateAttendance)
		auth.Put("/api/v1/attendances/{id}", controllers.UpdateAttendance)
		auth.Delete("/api/v1/attendances/{id}", controllers.DeleteAttendance)
		
		// Admin routes
		auth.Get("/api/v1/admin", controllers.GetAllPersonnel)
		//auth.Get("/api/v1/events/{course}", controllers.GetEventByCourse)
		
		auth.Put("/api/v1/admin/{username}", controllers.UpdateAdmin)
		auth.Delete("/api/v1/admin/{username}", controllers.DeleteAdmin)
		
		// Personnel routes
		auth.Get("/api/v1/personnels", controllers.GetAllPersonnel)
		auth.Get("/api/v1/personnels/{id}", controllers.GetPersonnelById)
		//auth.Get("/api/v1/events/{course}", controllers.GetEventByCourse)
		auth.Put("/api/v1/personnels/{id}", controllers.UpdatePersonnel)
		auth.Delete("/api/v1/personnels/{id}", controllers.DeletePersonnel)
	})
	return router
}
