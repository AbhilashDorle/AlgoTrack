package firebase

import (
	"context"
	"log"
	"path/filepath"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

/*
	This function initializes the firebase client. After this we can verify tokens

Everything in the function is referred from the firebase docs
*/

var AuthClient *auth.Client

func InitFirebase() {
	// opt := option.WithCredentialsFile("../../firebase/serviceAccountKey.json")
	// Define your relative path
	relativePath := "../../firebase/serviceAccountKey.json"

	// Convert to absolute path for debugging
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		log.Fatalf("Error resolving absolute path: %v", err)
	}

	log.Println("🔍 Loading Firebase credentials from:", absPath)

	opt := option.WithCredentialsFile(absPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing Firebase app: %v", err)
	}

	AuthClient, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v", err)
	}

	log.Println("✅ Firebase initialized successfully:", AuthClient != nil)
}
