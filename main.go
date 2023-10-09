package main

import (
	"boilerplate/internal/api"
	"boilerplate/internal/features/auth"
	"boilerplate/internal/features/example"
	"boilerplate/internal/features/user"
	"boilerplate/internal/pkg"
	"boilerplate/internal/shared/config"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.InitializeConfig()
	if err != nil {
		panic(err)
	}

	if c.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create data sources
	serviceAccountKeyPath, err := filepath.Abs("./firebaseServiceAccountKey.json")
	if err != nil {
		panic("[ERROR] Unable to load firebaseServiceAccountKey.json")
	}
	authClient, err := pkg.InitializeFirebase(serviceAccountKeyPath)

	db, err := pkg.InitializeSqlite(c.Database)
	if err != nil {
		panic(err)
	}

	// Create repositories
	authRepo := auth.NewAuthRepository(authClient)
	exampleRepo := example.NewExampleRepository(db)
	userRepo := user.NewUserRepository(db)

	// Create services
	authService := auth.NewAuthService(authRepo)
	exampleService := example.NewExampleService(exampleRepo)
	userService := user.NewUserService(userRepo)

	s := api.InitializeApi("", c.Server.Port, authService, exampleService, userService)

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
