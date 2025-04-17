package routes

import (
	"net/http"

	"cinema-booking-system/internal/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// InitRoutes initializes all API routes
func InitRoutes(router *mux.Router, db *gorm.DB) {
	// Create API handlers
	apiHandler := handlers.NewHandler(db)

	// Create template handler
	templateHandler, err := handlers.NewTemplateHandler()
	if err != nil {
		panic(err)
	}

	// Frontend routes
	router.HandleFunc("/", templateHandler.HomeHandler).Methods("GET")
	router.HandleFunc("/movies", templateHandler.MoviesHandler).Methods("GET")
	router.HandleFunc("/movies/{id:[0-9]+}", templateHandler.MovieDetailHandler).Methods("GET")
	router.HandleFunc("/booking", templateHandler.SeatBookingHandler).Methods("GET")
	router.HandleFunc("/login", templateHandler.LoginHandler).Methods("GET")
	router.HandleFunc("/register", templateHandler.RegisterHandler).Methods("GET")
	router.HandleFunc("/bookings", templateHandler.BookingsHandler).Methods("GET")

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	// Authentication routes
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", apiHandler.Register).Methods("POST")
	auth.HandleFunc("/login", apiHandler.Login).Methods("POST")

	// Movie routes
	movies := api.PathPrefix("/movies").Subrouter()
	movies.HandleFunc("", apiHandler.GetMovies).Methods("GET")
	movies.HandleFunc("/{id}", apiHandler.GetMovie).Methods("GET")

	// Show routes
	shows := api.PathPrefix("/shows").Subrouter()
	shows.HandleFunc("", apiHandler.GetShows).Methods("GET")
	shows.HandleFunc("/{id}", apiHandler.GetShow).Methods("GET")
	shows.HandleFunc("/{id}/seats", apiHandler.GetShowSeats).Methods("GET")

	// Booking routes - protected by authentication middleware
	// bookings := api.PathPrefix("/bookings").Subrouter()
	// bookings.Use(middleware.JWTAuth)
	// bookings.HandleFunc("", apiHandler.GetUserBookings).Methods("GET")
	// bookings.HandleFunc("", apiHandler.CreateBooking).Methods("POST")

	// Serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
}
