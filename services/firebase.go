package services

import (
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"golang.org/x/net/context"

	"google.golang.org/api/option"
)

// FirebaseService contains methods for finding and manipulating resources
type FirebaseService struct {
	firebase *firebase.App
}

// NewFirebaseService creates an instance of FirebaseService
func NewFirebaseService(firebaseCredentialsFilePath string) *FirebaseService {
	opt := option.WithCredentialsFile(firebaseCredentialsFilePath)
	firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return &FirebaseService{firebase: firebaseApp}
}

// Verify verify an id token
func (fs FirebaseService) Verify(id string) (*auth.Token, error) {
	var err error
	client, err := fs.firebase.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
		return nil, err
	}

	token, err := client.VerifyIDToken(id)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
		return nil, err
	}

	return token, err
}

// GetUser get a user by uid
func (fs FirebaseService) GetUser(uid string) *auth.UserRecord {
	// Get an auth client from the firebase.App
	ctx := context.Background()
	client, err := fs.firebase.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	user, err := client.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	log.Printf("Successfully fetched user data: %v\n", user)

	return user
}
