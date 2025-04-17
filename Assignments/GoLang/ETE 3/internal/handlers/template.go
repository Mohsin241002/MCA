package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// TemplateHandler handles rendering of HTML templates
type TemplateHandler struct {
	templates *template.Template
}

// NewTemplateHandler creates a new template handler
func NewTemplateHandler() *TemplateHandler {
	// Parse all templates in the templates directory
	templates, err := template.ParseGlob(filepath.Join("web", "templates", "*.html"))
	if err != nil {
		// In production, handle this more gracefully
		panic(err)
	}

	return &TemplateHandler{
		templates: templates,
	}
}

// HomeHandler renders the home page
func (h *TemplateHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "home.html", nil)
}

// MoviesHandler renders the movies page
func (h *TemplateHandler) MoviesHandler(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "movies.html", nil)
}

// MovieDetailHandler renders the movie detail page
func (h *TemplateHandler) MovieDetailHandler(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "movie_detail.html", nil)
}

// SeatBookingHandler renders the seat booking page
func (h *TemplateHandler) SeatBookingHandler(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "seat_booking.html", nil)
}

// LoginHandler renders the login page
func (h *TemplateHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "login.html", nil)
}

// RegisterHandler renders the register page
func (h *TemplateHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "register.html", nil)
}

// BookingsHandler renders the user bookings page
func (h *TemplateHandler) BookingsHandler(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "bookings.html", nil)
}
