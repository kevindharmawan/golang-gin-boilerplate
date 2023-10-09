package pkg

import (
	"context"
	"errors"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func InitializeFirebase(serviceAccountKeyPath string) (*auth.Client, error) {
	var opt option.ClientOption

	// TODO: Refactor
	if _, err := os.Stat(serviceAccountKeyPath); errors.Is(err, os.ErrNotExist) {
		fmt.Println("firebaseServiceAccountKey.json doesn't exists, use environment instead")
		opt = option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_CREDENTIALS")))
	} else {
		opt = option.WithCredentialsFile(serviceAccountKeyPath)
	}

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return auth, nil
}
