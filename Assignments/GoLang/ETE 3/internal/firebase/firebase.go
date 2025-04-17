package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

// FirebaseApp contains Firebase services
type FirebaseApp struct {
	App       *firebase.App
	Auth      *auth.Client
	Firestore *firestore.Client
}

// InitFirebase initializes Firebase services
func InitFirebase() (*FirebaseApp, error) {
	ctx := context.Background()
	var app *firebase.App
	var err error

	// For development/demo purposes, we'll skip Firebase authentication
	// In a real app, you would use proper credentials
	log.Println("Initializing Firebase in demo mode...")

	// Initialize with application default credentials
	config := &firebase.Config{
		ProjectID: "cinema-34494",
	}

	// If running locally without proper credentials, this will use
	// the Firebase emulator if available
	app, err = firebase.NewApp(ctx, config)
	if err != nil {
		// If that fails, create a mock implementation
		log.Printf("Warning: Unable to initialize Firebase normally: %v", err)
		log.Println("Continuing in full demo mode with limited functionality")

		// In a real application, you should not proceed without proper Firebase setup
		// For demo purposes only:
		return &FirebaseApp{
			App:       nil,
			Auth:      nil,
			Firestore: nil,
		}, nil
	}

	// Initialize Firebase Authentication
	auth, err := app.Auth(ctx)
	if err != nil {
		log.Printf("Warning: Error initializing Firebase Auth: %v", err)
		auth = nil
	}

	// Initialize Firebase Firestore
	firestore, err := app.Firestore(ctx)
	if err != nil {
		log.Printf("Warning: Error initializing Firestore: %v", err)
		firestore = nil
	}

	fbApp := &FirebaseApp{
		App:       app,
		Auth:      auth,
		Firestore: firestore,
	}

	if fbApp.Auth == nil && fbApp.Firestore == nil {
		log.Println("Warning: Running with no Firebase services available!")
		log.Println("This is for demonstration purposes only.")
	}

	return fbApp, nil
}
