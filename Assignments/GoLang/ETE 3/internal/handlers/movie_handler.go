package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"cinema-booking-system/internal/models"
	"cinema-booking-system/internal/services"

	"github.com/gorilla/mux"
)

// MovieHandler handles HTTP requests related to movies
type MovieHandler struct {
	movieService *services.MovieService
}

// NewMovieHandler creates a new movie handler
func NewMovieHandler(movieService *services.MovieService) *MovieHandler {
	return &MovieHandler{
		movieService: movieService,
	}
}

// CreateMovie creates a new movie
func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set created and updated timestamps
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()

	err = h.movieService.CreateMovie(&movie)
	if err != nil {
		http.Error(w, "Failed to create movie: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Movie created successfully",
		"id":      movie.ID,
		"movie":   movie,
	})
}

// GetMovie retrieves a movie by ID
func (h *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	movie, err := h.movieService.GetMovieByID(id)
	if err != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

// UpdateMovie updates a movie
func (h *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Get existing movie
	existingMovie, err := h.movieService.GetMovieByID(id)
	if err != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	// Decode updates
	var movieUpdates models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movieUpdates); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update fields (only non-empty ones)
	if movieUpdates.Title != "" {
		existingMovie.Title = movieUpdates.Title
	}
	if movieUpdates.Description != "" {
		existingMovie.Description = movieUpdates.Description
	}
	if movieUpdates.Duration != 0 {
		existingMovie.Duration = movieUpdates.Duration
	}
	if movieUpdates.Genre != "" {
		existingMovie.Genre = movieUpdates.Genre
	}
	if movieUpdates.Language != "" {
		existingMovie.Language = movieUpdates.Language
	}
	if !movieUpdates.ReleaseDate.IsZero() {
		existingMovie.ReleaseDate = movieUpdates.ReleaseDate
	}
	if movieUpdates.Director != "" {
		existingMovie.Director = movieUpdates.Director
	}
	if movieUpdates.Cast != "" {
		existingMovie.Cast = movieUpdates.Cast
	}
	if movieUpdates.PosterURL != "" {
		existingMovie.PosterURL = movieUpdates.PosterURL
	}

	// Update timestamp
	existingMovie.UpdatedAt = time.Now()

	// Save the updated movie
	if err := h.movieService.UpdateMovie(existingMovie); err != nil {
		http.Error(w, "Failed to update movie: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Movie updated successfully",
	})
}

// DeleteMovie deletes a movie
func (h *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.movieService.DeleteMovie(id); err != nil {
		http.Error(w, "Failed to delete movie: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Movie deleted successfully",
	})
}

// ListMovies lists all movies
func (h *MovieHandler) ListMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.movieService.ListMovies()
	if err != nil {
		http.Error(w, "Failed to list movies: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
