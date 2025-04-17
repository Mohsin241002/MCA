package models

import (
	"time"
)

// User represents a user of the system
type User struct {
	ID          string    `json:"id" firestore:"id"`
	Username    string    `json:"username" firestore:"username"`
	Email       string    `json:"email" firestore:"email"`
	FullName    string    `json:"full_name" firestore:"full_name"`
	PhoneNumber string    `json:"phone_number" firestore:"phone_number"`
	Role        string    `json:"role" firestore:"role"`
	CreatedAt   time.Time `json:"created_at" firestore:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" firestore:"updated_at"`
}

// Movie represents a movie in the cinema
type Movie struct {
	ID          string    `json:"id" firestore:"id"`
	Title       string    `json:"title" firestore:"title"`
	Description string    `json:"description" firestore:"description"`
	Duration    int       `json:"duration" firestore:"duration"` // Duration in minutes
	Genre       string    `json:"genre" firestore:"genre"`
	Language    string    `json:"language" firestore:"language"`
	ReleaseDate time.Time `json:"release_date" firestore:"release_date"`
	Director    string    `json:"director" firestore:"director"`
	Cast        string    `json:"cast" firestore:"cast"`
	PosterURL   string    `json:"poster_url" firestore:"poster_url"`
	CreatedAt   time.Time `json:"created_at" firestore:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" firestore:"updated_at"`
}

// Show represents a movie screening
type Show struct {
	ID         string    `json:"id" firestore:"id"`
	MovieID    string    `json:"movie_id" firestore:"movie_id"`
	ShowTime   time.Time `json:"show_time" firestore:"show_time"`
	EndTime    time.Time `json:"end_time" firestore:"end_time"`
	HallNumber int       `json:"hall_number" firestore:"hall_number"`
	Price      float64   `json:"price" firestore:"price"`
	CreatedAt  time.Time `json:"created_at" firestore:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" firestore:"updated_at"`
}

// Seat represents a seat in a cinema hall
type Seat struct {
	ID         string    `json:"id" firestore:"id"`
	ShowID     string    `json:"show_id" firestore:"show_id"`
	SeatNumber string    `json:"seat_number" firestore:"seat_number"` // e.g., "A1", "B2"
	IsBooked   bool      `json:"is_booked" firestore:"is_booked"`
	CreatedAt  time.Time `json:"created_at" firestore:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" firestore:"updated_at"`
}

// Booking represents a booking in the system
type Booking struct {
	ID            string    `json:"id" firestore:"id"`
	UserID        string    `json:"user_id" firestore:"user_id"`
	ShowID        string    `json:"show_id" firestore:"show_id"`
	SeatID        string    `json:"seat_id" firestore:"seat_id"`
	BookingDate   time.Time `json:"booking_date" firestore:"booking_date"`
	TotalAmount   float64   `json:"total_amount" firestore:"total_amount"`
	PaymentStatus string    `json:"payment_status" firestore:"payment_status"`
	BookingStatus string    `json:"booking_status" firestore:"booking_status"`
	CreatedAt     time.Time `json:"created_at" firestore:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" firestore:"updated_at"`
}
