package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"cinema-booking-system/internal/models"
	"cinema-booking-system/internal/services"

	"github.com/gorilla/mux"
)

// BookingHandler handles HTTP requests related to bookings
type BookingHandler struct {
	bookingService *services.BookingService
}

// NewBookingHandler creates a new booking handler
func NewBookingHandler(bookingService *services.BookingService) *BookingHandler {
	return &BookingHandler{
		bookingService: bookingService,
	}
}

// CreateBooking creates a new booking
func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set created and updated timestamps
	booking.CreatedAt = time.Now()
	booking.UpdatedAt = time.Now()

	err = h.bookingService.CreateBooking(&booking)
	if err != nil {
		http.Error(w, "Failed to create booking: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Booking created successfully",
		"id":      booking.ID,
	})
}

// GetBooking retrieves a booking by ID
func (h *BookingHandler) GetBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	booking, err := h.bookingService.GetBookingByID(id)
	if err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

// CancelBooking cancels a booking
func (h *BookingHandler) CancelBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.bookingService.CancelBooking(id); err != nil {
		http.Error(w, "Failed to cancel booking: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Booking cancelled successfully",
	})
}

// GetUserBookings retrieves all bookings for a user
func (h *BookingHandler) GetUserBookings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	bookings, err := h.bookingService.GetUserBookings(userID)
	if err != nil {
		http.Error(w, "Failed to get user bookings: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

// ListBookings lists all bookings
func (h *BookingHandler) ListBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := h.bookingService.ListBookings()
	if err != nil {
		http.Error(w, "Failed to list bookings: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}
