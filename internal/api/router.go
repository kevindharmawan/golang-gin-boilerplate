package api

import (
	"boilerplate/internal/api/middleware"
	"boilerplate/internal/features/example"
	"boilerplate/internal/features/user"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(
	corsMiddleware *middleware.CorsMiddleware,
	authMiddleware *middleware.AuthMiddleware,
	exampleHandler *example.ExampleHandler,
	userHandler *user.UserHandler,
) *gin.Engine {
	app := gin.Default()

	apiRoute := app.Group("/api")
	apiRoute.Use(corsMiddleware.CorsMiddleware)
	apiRoute.Use(authMiddleware.AuthMiddleware)
	{
		apiRoute.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	exampleRoute := apiRoute.Group("/example")
	{
		exampleRoute.GET("/get", exampleHandler.GetAllExample)
		exampleRoute.GET("/get/:id", exampleHandler.GetExampleById)
		exampleRoute.POST("/create", exampleHandler.CreateExample)
		exampleRoute.PATCH("/update/:id", exampleHandler.UpdateExample)
		exampleRoute.DELETE("/delete/:id", exampleHandler.DeleteExample)
	}

	userRoute := apiRoute.Group("/user")
	{
		userRoute.GET("/get", authMiddleware.UserRequiredMiddleware, userHandler.GetCurrentUser)
		userRoute.POST("/create", userHandler.CreateUser)
		userRoute.PATCH("/update", authMiddleware.UserRequiredMiddleware, userHandler.UpdateCurrentUser)
		userRoute.DELETE("/delete", authMiddleware.UserRequiredMiddleware, userHandler.DeleteCurrentUser)
	}

	return app
}
