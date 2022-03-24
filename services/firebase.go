package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var FirebaseAuth *auth.Client

func InitializeFirebase() {
	serviceAccountKeyPath, err := filepath.Abs("./firebaseServiceAccountKey.json")
	if err != nil {
		panic("[ERROR] Unable to load firebaseServiceAccountKey.json")
	}

	var opt option.ClientOption

	if _, err := os.Stat(serviceAccountKeyPath); errors.Is(err, os.ErrNotExist) {
		fmt.Println("firebaseServiceAccountKey.json doesn't exists, use environment instead")
		opt = option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_CREDENTIALS")))
	} else {
		opt = option.WithCredentialsFile(serviceAccountKeyPath)
	}

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("[ERROR] Firebase load error")
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		panic("[ERROR] Firebase load error")
	}

	FirebaseAuth = auth
}
