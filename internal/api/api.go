package api

import (
	"boilerplate/internal/api/middleware"
	"boilerplate/internal/features/auth"
	"boilerplate/internal/features/example"
	"boilerplate/internal/features/user"
	"fmt"
	"net/http"
	"time"
)

func InitializeApi(
	host string,
	port int,
	authService auth.AuthService,
	exampleService example.ExampleService,
	userService user.UserService,
) *http.Server {
	corsMiddleware := middleware.NewCorsMiddleware()
	authMiddleware := middleware.NewAuthMiddleware(authService, userService)

	exampleHandler := example.NewExampleHandler(exampleService)
	userHandler := user.NewUserHandler(userService)

	return &http.Server{
		Addr: fmt.Sprintf("%v:%d", host, port),
		Handler: InitializeRouter(
			corsMiddleware,
			authMiddleware,
			exampleHandler,
			userHandler,
		),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
