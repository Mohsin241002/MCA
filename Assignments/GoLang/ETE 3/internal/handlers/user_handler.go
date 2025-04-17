package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"cinema-booking-system/internal/models"
	"cinema-booking-system/internal/services"

	"github.com/gorilla/mux"
)

// UserHandler handles HTTP requests related to users
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Register handles user registration
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set created and updated timestamps
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Default role to "user"
	if user.Role == "" {
		user.Role = "user"
	}

	err = h.userService.Register(&user)
	if err != nil {
		http.Error(w, "Failed to register user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
		"id":      user.ID,
	})
}

// Login handles user authentication
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// We'll rely on Firebase Auth for actual authentication
	// This is a simplified example

	user, err := h.userService.GetUserByEmail(credentials.Email)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// In a real application, we would validate the password and generate a JWT token
	// For simplicity, we'll just return the user ID

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
		"id":      user.ID,
	})
}

// GetUser retrieves a user by ID
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUser updates a user
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var userUpdates models.User
	if err := json.NewDecoder(r.Body).Decode(&userUpdates); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get existing user
	existingUser, err := h.userService.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Update fields (only non-empty ones)
	if userUpdates.Username != "" {
		existingUser.Username = userUpdates.Username
	}
	if userUpdates.Email != "" {
		existingUser.Email = userUpdates.Email
	}
	if userUpdates.FullName != "" {
		existingUser.FullName = userUpdates.FullName
	}
	if userUpdates.PhoneNumber != "" {
		existingUser.PhoneNumber = userUpdates.PhoneNumber
	}

	// Update the user
	existingUser.UpdatedAt = time.Now()

	if err := h.userService.UpdateUser(existingUser); err != nil {
		http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User updated successfully",
	})
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.userService.DeleteUser(id); err != nil {
		http.Error(w, "Failed to delete user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User deleted successfully",
	})
}

// ListUsers lists all users
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.ListUsers()
	if err != nil {
		http.Error(w, "Failed to list users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
