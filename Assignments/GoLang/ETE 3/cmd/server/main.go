package main

import (
	"log"
	"net/http"

	"cinema-booking-system/internal/firebase"
	"cinema-booking-system/internal/handlers"
	"cinema-booking-system/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting Cinema Booking System...")

	// Initialize Firebase
	fbApp, err := firebase.InitFirebase()
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}

	// Initialize services with Firebase
	userService := services.NewUserService(fbApp)
	movieService := services.NewMovieService(fbApp)
	bookingService := services.NewBookingService(fbApp)

	// Initialize handlers with services
	userHandler := handlers.NewUserHandler(userService)
	movieHandler := handlers.NewMovieHandler(movieService)
	bookingHandler := handlers.NewBookingHandler(bookingService)
	templateHandler := handlers.NewTemplateHandler()

	// Create router
	router := mux.NewRouter()

	// API routes
	apiRouter := router.PathPrefix("/api").Subrouter()

	// User routes
	apiRouter.HandleFunc("/users/register", userHandler.Register).Methods("POST")
	apiRouter.HandleFunc("/users/login", userHandler.Login).Methods("POST")
	apiRouter.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	apiRouter.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	apiRouter.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	apiRouter.HandleFunc("/users", userHandler.ListUsers).Methods("GET")

	// Movie routes
	apiRouter.HandleFunc("/movies", movieHandler.CreateMovie).Methods("POST")
	apiRouter.HandleFunc("/movies/{id}", movieHandler.GetMovie).Methods("GET")
	apiRouter.HandleFunc("/movies/{id}", movieHandler.UpdateMovie).Methods("PUT")
	apiRouter.HandleFunc("/movies/{id}", movieHandler.DeleteMovie).Methods("DELETE")
	apiRouter.HandleFunc("/movies", movieHandler.ListMovies).Methods("GET")

	// Booking routes
	apiRouter.HandleFunc("/bookings", bookingHandler.CreateBooking).Methods("POST")
	apiRouter.HandleFunc("/bookings/{id}", bookingHandler.GetBooking).Methods("GET")
	apiRouter.HandleFunc("/bookings/{id}/cancel", bookingHandler.CancelBooking).Methods("PUT")
	apiRouter.HandleFunc("/bookings/user/{userId}", bookingHandler.GetUserBookings).Methods("GET")
	apiRouter.HandleFunc("/bookings", bookingHandler.ListBookings).Methods("GET")

	// Frontend routes
	router.HandleFunc("/", templateHandler.HomeHandler).Methods("GET")
	router.HandleFunc("/movies", templateHandler.MoviesHandler).Methods("GET")
	router.HandleFunc("/movies/{id:[0-9]+}", templateHandler.MovieDetailHandler).Methods("GET")
	router.HandleFunc("/booking", templateHandler.SeatBookingHandler).Methods("GET")
	router.HandleFunc("/login", templateHandler.LoginHandler).Methods("GET")
	router.HandleFunc("/register", templateHandler.RegisterHandler).Methods("GET")
	router.HandleFunc("/bookings", templateHandler.BookingsHandler).Methods("GET")

	// Serve static files
	fs := http.FileServer(http.Dir("./web/static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Start server
	log.Println("Server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
