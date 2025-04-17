package services

import (
	"context"
	"errors"
	"log"
	"time"

	"cinema-booking-system/internal/firebase"
	"cinema-booking-system/internal/models"

	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/iterator"
)

// UserService handles user-related operations
type UserService struct {
	fbApp *firebase.FirebaseApp
	// In-memory users for demo when Firebase is not available
	demoUsers map[string]*models.User
}

// NewUserService creates a new user service
func NewUserService(fbApp *firebase.FirebaseApp) *UserService {
	return &UserService{
		fbApp:     fbApp,
		demoUsers: make(map[string]*models.User),
	}
}

// Register creates a new user in the system
func (s *UserService) Register(user *models.User) error {
	// For demo mode without Firebase
	if s.fbApp.Auth == nil || s.fbApp.Firestore == nil {
		log.Println("Running in demo mode: Creating user in memory")
		// For demo, generate a simple ID
		if user.ID == "" {
			user.ID = "user_" + time.Now().Format("20060102150405")
		}
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		s.demoUsers[user.ID] = user
		return nil
	}

	ctx := context.Background()

	// Create a user in Firebase Authentication
	params := (&auth.UserToCreate{}).
		Email(user.Email).
		Password("password"). // You should pass this in securely
		DisplayName(user.FullName)

	userRecord, err := s.fbApp.Auth.CreateUser(ctx, params)
	if err != nil {
		log.Printf("Error creating auth user: %v", err)
		return err
	}

	// Set user ID from the Auth UID
	user.ID = userRecord.UID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Save user to Firestore
	_, err = s.fbApp.Firestore.Collection("users").Doc(user.ID).Set(ctx, user)
	if err != nil {
		log.Printf("Error saving user to Firestore: %v", err)
		return err
	}

	return nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id string) (*models.User, error) {
	// For demo mode without Firebase
	if s.fbApp.Firestore == nil {
		log.Println("Running in demo mode: Getting user from memory")
		user, exists := s.demoUsers[id]
		if !exists {
			return nil, errors.New("user not found")
		}
		return user, nil
	}

	ctx := context.Background()

	doc, err := s.fbApp.Firestore.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByEmail retrieves a user by email
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	// For demo mode without Firebase
	if s.fbApp.Firestore == nil {
		log.Println("Running in demo mode: Getting user from memory")
		// Search in memory
		for _, user := range s.demoUsers {
			if user.Email == email {
				return user, nil
			}
		}
		return nil, errors.New("user not found")
	}

	ctx := context.Background()

	query := s.fbApp.Firestore.Collection("users").Where("email", "==", email).Limit(1)
	iter := query.Documents(ctx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err == iterator.Done {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates a user in the system
func (s *UserService) UpdateUser(user *models.User) error {
	// For demo mode without Firebase
	if s.fbApp.Firestore == nil {
		log.Println("Running in demo mode: Updating user in memory")
		_, exists := s.demoUsers[user.ID]
		if !exists {
			return errors.New("user not found")
		}
		user.UpdatedAt = time.Now()
		s.demoUsers[user.ID] = user
		return nil
	}

	ctx := context.Background()

	user.UpdatedAt = time.Now()

	_, err := s.fbApp.Firestore.Collection("users").Doc(user.ID).Set(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes a user from the system
func (s *UserService) DeleteUser(id string) error {
	// For demo mode without Firebase
	if s.fbApp.Firestore == nil || s.fbApp.Auth == nil {
		log.Println("Running in demo mode: Deleting user from memory")
		delete(s.demoUsers, id)
		return nil
	}

	ctx := context.Background()

	// Delete from Firestore
	_, err := s.fbApp.Firestore.Collection("users").Doc(id).Delete(ctx)
	if err != nil {
		return err
	}

	// Delete from Authentication
	if err := s.fbApp.Auth.DeleteUser(ctx, id); err != nil {
		return err
	}

	return nil
}

// ListUsers retrieves all users from the system
func (s *UserService) ListUsers() ([]*models.User, error) {
	// For demo mode without Firebase
	if s.fbApp.Firestore == nil {
		log.Println("Running in demo mode: Listing users from memory")
		users := make([]*models.User, 0, len(s.demoUsers))
		for _, user := range s.demoUsers {
			users = append(users, user)
		}
		return users, nil
	}

	ctx := context.Background()

	var users []*models.User

	iter := s.fbApp.Firestore.Collection("users").Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var user models.User
		if err := doc.DataTo(&user); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}
