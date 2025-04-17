package services

import (
	"context"
	"errors"
	"log"
	"time"

	"cinema-booking-system/internal/firebase"
	"cinema-booking-system/internal/models"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// BookingService handles booking-related operations
type BookingService struct {
	fbApp *firebase.FirebaseApp
	// In-memory data for demo when Firebase is not available
	demoBookings map[string]*models.Booking
	demoSeats    map[string]*models.Seat
}

// NewBookingService creates a new booking service
func NewBookingService(fbApp *firebase.FirebaseApp) *BookingService {
	return &BookingService{
		fbApp:        fbApp,
		demoBookings: make(map[string]*models.Booking),
		demoSeats:    make(map[string]*models.Seat),
	}
}

// CreateBooking creates a new booking in the system
func (s *BookingService) CreateBooking(booking *models.Booking) error {
	// For demo mode without Firebase
	if s.fbApp.Firestore == nil {
		log.Println("Running in demo mode: Creating booking in memory")

		// Check if seat exists and is available
		seat, exists := s.demoSeats[booking.SeatID]
		if !exists {
			// For demo, create the seat if it doesn't exist
			seat = &models.Seat{
				ID:         booking.SeatID,
				ShowID:     booking.ShowID,
				SeatNumber: "A1", // Default for demo
				IsBooked:   false,
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			}
		}

		if seat.IsBooked {
			return errors.New("seat is already booked")
		}

		// Generate booking ID
		booking.ID = "booking_" + time.Now().Format("20060102150405")
		booking.BookingDate = time.Now()
		booking.CreatedAt = time.Now()
		booking.UpdatedAt = time.Now()
		booking.BookingStatus = "confirmed"

		// Mark seat as booked
		seat.IsBooked = true
		seat.UpdatedAt = time.Now()

		// Save both to our in-memory store
		s.demoBookings[booking.ID] = booking
		s.demoSeats[seat.ID] = seat

		return nil
	}

	ctx := context.Background()

	// Start a transaction to ensure all operations succeed or fail together
	err := s.fbApp.Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		// First check if the seat is available
		seatRef := s.fbApp.Firestore.Collection("seats").Doc(booking.SeatID)
		seat, err := tx.Get(seatRef)
		if err != nil {
			return err
		}

		var seatData models.Seat
		if err := seat.DataTo(&seatData); err != nil {
			return err
		}

		if seatData.IsBooked {
			return errors.New("seat is already booked")
		}

		// Generate a booking ID
		bookingRef := s.fbApp.Firestore.Collection("bookings").NewDoc()
		booking.ID = bookingRef.ID
		booking.BookingDate = time.Now()
		booking.CreatedAt = time.Now()
		booking.UpdatedAt = time.Now()
		booking.BookingStatus = "confirmed"

		// Update the seat to be booked
		seatData.IsBooked = true
		seatData.UpdatedAt = time.Now()

		// Create the booking and update the seat in the transaction
		if err := tx.Set(bookingRef, booking); err != nil {
			return err
		}

		if err := tx.Set(seatRef, seatData); err != nil {
			return err
		}

		return nil
	})

	return err
}

// GetBookingByID retrieves a booking by ID
func (s *BookingService) GetBookingByID(id string) (*models.Booking, error) {
	// For demo mode without Firebase
	if s.fbApp.Firestore == nil {
		log.Println("Running in demo mode: Getting booking from memory")
		booking, exists := s.demoBookings[id]
		if !exists {
			return nil, errors.New("booking not found")
		}
		return booking, nil
	}

	ctx := context.Background()

	doc, err := s.fbApp.Firestore.Collection("bookings").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var booking models.Booking
	if err := doc.DataTo(&booking); err != nil {
		return nil, err
	}

	return &booking, nil
}

// CancelBooking cancels a booking
func (s *BookingService) CancelBooking(id string) error {
	// For demo mode without Firebase
	if s.fbApp.Firestore == nil {
		log.Println("Running in demo mode: Cancelling booking in memory")

		booking, exists := s.demoBookings[id]
		if !exists {
			return errors.New("booking not found")
		}

		seat, exists := s.demoSeats[booking.SeatID]
		if !exists {
			return errors.New("seat not found")
		}

		// Update booking
		booking.BookingStatus = "cancelled"
		booking.UpdatedAt = time.Now()

		// Free up seat
		seat.IsBooked = false
		seat.UpdatedAt = time.Now()

		// Update in-memory storage
		s.demoBookings[id] = booking
		s.demoSeats[seat.ID] = seat

		return nil
	}

	ctx := context.Background()

	err := s.fbApp.Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		// Get the booking
		bookingRef := s.fbApp.Firestore.Collection("bookings").Doc(id)
		booking, err := tx.Get(bookingRef)
		if err != nil {
			return err
		}

		var bookingData models.Booking
		if err := booking.DataTo(&bookingData); err != nil {
			return err
		}

		// Get the seat
		seatRef := s.fbApp.Firestore.Collection("seats").Doc(bookingData.SeatID)
		seat, err := tx.Get(seatRef)
		if err != nil {
			return err
		}

		var seatData models.Seat
		if err := seat.DataTo(&seatData); err != nil {
			return err
		}

		// Update booking status
		bookingData.BookingStatus = "cancelled"
		bookingData.UpdatedAt = time.Now()

		// Free up the seat
		seatData.IsBooked = false
		seatData.UpdatedAt = time.Now()

		// Update both in the transaction
		if err := tx.Set(bookingRef, bookingData); err != nil {
			return err
		}

		if err := tx.Set(seatRef, seatData); err != nil {
			return err
		}

		return nil
	})

	return err
}

// GetUserBookings gets all bookings for a user
func (s *BookingService) GetUserBookings(userID string) ([]*models.Booking, error) {
	// For demo mode without Firebase
	if s.fbApp.Firestore == nil {
		log.Println("Running in demo mode: Getting user bookings from memory")
		var bookings []*models.Booking

		for _, booking := range s.demoBookings {
			if booking.UserID == userID {
				bookings = append(bookings, booking)
			}
		}

		return bookings, nil
	}

	ctx := context.Background()

	var bookings []*models.Booking

	query := s.fbApp.Firestore.Collection("bookings").Where("user_id", "==", userID)
	iter := query.Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var booking models.Booking
		if err := doc.DataTo(&booking); err != nil {
			return nil, err
		}

		bookings = append(bookings, &booking)
	}

	return bookings, nil
}

// ListBookings retrieves all bookings from the system
func (s *BookingService) ListBookings() ([]*models.Booking, error) {
	// For demo mode without Firebase
	if s.fbApp.Firestore == nil {
		log.Println("Running in demo mode: Listing bookings from memory")
		bookings := make([]*models.Booking, 0, len(s.demoBookings))
		for _, booking := range s.demoBookings {
			bookings = append(bookings, booking)
		}
		return bookings, nil
	}

	ctx := context.Background()

	var bookings []*models.Booking

	iter := s.fbApp.Firestore.Collection("bookings").Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var booking models.Booking
		if err := doc.DataTo(&booking); err != nil {
			return nil, err
		}

		bookings = append(bookings, &booking)
	}

	return bookings, nil
}
