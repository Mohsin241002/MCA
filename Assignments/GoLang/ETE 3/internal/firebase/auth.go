// This file is no longer needed - we're using firebase.google.com/go/v4/auth directly

package firebase

// UserToCreate holds the necessary information for creating a new user in Firebase Authentication
type UserToCreate struct {
	Email         string
	Password      string
	DisplayName   string
	PhoneNumber   string
	PhotoURL      string
	Disabled      bool
	EmailVerified bool
}
