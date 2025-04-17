package services

import (
	"context"
	"errors"
	"log"
	"time"

	"cinema-booking-system/internal/firebase"
	"cinema-booking-system/internal/models"

	"google.golang.org/api/iterator"
)

// MovieService handles movie-related operations
type MovieService struct {
	fbApp *firebase.FirebaseApp
	// In-memory movies for demo when Firebase is not available
	demoMovies map[string]*models.Movie
}

// NewMovieService creates a new movie service
func NewMovieService(fbApp *firebase.FirebaseApp) *MovieService {
	return &MovieService{
		fbApp:      fbApp,
		demoMovies: make(map[string]*models.Movie),
	}
}

// CreateMovie creates a new movie in the system
func (s *MovieService) CreateMovie(movie *models.Movie) error {
	// For demo mode without Firebase
	if s.fbApp.Firestore == nil {
		log.Println("Running in demo mode: Creating movie in memory")
		// For demo, generate a simple ID
		if movie.ID == "" {
			movie.ID = "movie_" + time.Now().Format("20060102150405")
		}
		movie.CreatedAt = time.Now()
		movie.UpdatedAt = time.Now()
		s.demoMovies[movie.ID] = movie
		return nil
	}

	ctx := context.Background()

	// Generate a new document ID
	ref := s.fbApp.Firestore.Collection("movies").NewDoc()
	movie.ID = ref.ID
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()

	_, err := ref.Set(ctx, movie)
	if err != nil {
		return err
	}

	return nil
}

// GetMovieByID retrieves a movie by ID
func (s *MovieService) GetMovieByID(id string) (*models.Movie, error) {
	// For demo mode without Firebase
	if s.fbApp.Firestore == nil {
		log.Println("Running in demo mode: Getting movie from memory")
		movie, exists := s.demoMovies[id]
		if !exists {
			return nil, errors.New("movie not found")
		}
		return movie, nil
	}

	ctx := context.Background()

	doc, err := s.fbApp.Firestore.Collection("movies").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var movie models.Movie
	if err := doc.DataTo(&movie); err != nil {
		return nil, err
	}

	return &movie, nil
}

// UpdateMovie updates a movie in the system
func (s *MovieService) UpdateMovie(movie *models.Movie) error {
	// For demo mode without Firebase
	if s.fbApp.Firestore == nil {
		log.Println("Running in demo mode: Updating movie in memory")
		_, exists := s.demoMovies[movie.ID]
		if !exists {
			return errors.New("movie not found")
		}
		movie.UpdatedAt = time.Now()
		s.demoMovies[movie.ID] = movie
		return nil
	}

	ctx := context.Background()

	movie.UpdatedAt = time.Now()

	_, err := s.fbApp.Firestore.Collection("movies").Doc(movie.ID).Set(ctx, movie)
	if err != nil {
		return err
	}

	return nil
}

// DeleteMovie deletes a movie from the system
func (s *MovieService) DeleteMovie(id string) error {
	// For demo mode without Firebase
	if s.fbApp.Firestore == nil {
		log.Println("Running in demo mode: Deleting movie from memory")
		delete(s.demoMovies, id)
		return nil
	}

	ctx := context.Background()

	_, err := s.fbApp.Firestore.Collection("movies").Doc(id).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

// ListMovies retrieves all movies from the system
func (s *MovieService) ListMovies() ([]*models.Movie, error) {
	// For demo mode without Firebase
	if s.fbApp.Firestore == nil {
		log.Println("Running in demo mode: Listing movies from memory")
		movies := make([]*models.Movie, 0, len(s.demoMovies))
		for _, movie := range s.demoMovies {
			movies = append(movies, movie)
		}
		return movies, nil
	}

	ctx := context.Background()

	var movies []*models.Movie

	iter := s.fbApp.Firestore.Collection("movies").Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var movie models.Movie
		if err := doc.DataTo(&movie); err != nil {
			return nil, err
		}

		movies = append(movies, &movie)
	}

	return movies, nil
}
