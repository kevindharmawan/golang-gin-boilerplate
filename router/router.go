package router

import (
	"github.com/gin-gonic/gin"

	"boilerplate/controller/api"
	"boilerplate/controller/middleware"
	"boilerplate/utils"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()

	apiRoute := router.Group("/api")
	apiRoute.Use(
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
	)
	{
		profile := apiRoute.Group("/profile")
		{
			profile.GET("/get", utils.AuthOnly, api.GetMyProfile)
			profile.POST("/create", utils.AuthOnly, api.CreateProfile)
			profile.PATCH("/update", utils.AuthOnly, api.UpdateProfile)
			profile.DELETE("/delete", utils.AuthOnly, api.DeleteProfile)
		}
		example := apiRoute.Group("/example")
		{
			example.GET("/get", api.GetAllExample)
			example.GET("/get/:id", api.GetExampleById)
			example.POST("/create", api.CreateExample)
			example.PATCH("/update/:id", api.UpdateExample)
			example.DELETE("/delete/:id", api.DeleteExample)
		}
	}

	return router
}
