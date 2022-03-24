package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"boilerplate/constants"
	"boilerplate/models"
	"boilerplate/services"
)

var requireVerify = false

func AuthMiddleware(c *gin.Context) {
	firebaseAuth := services.FirebaseAuth

	authorizationToken := c.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

	if idToken == "" {
		c.Set(constants.IsAuthenticatedKey, false)
		c.Next()
		return
	}

	//verify token
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: "invalid token"})
		c.Abort()
		return
	}

	if requireVerify && !token.Claims["email_verified"].(bool) {
		c.JSON(http.StatusBadRequest, models.Response{Error: "account not verified"})
		c.Abort()
		return
	}

	c.Set(constants.UserIDKey, token.UID)
	c.Next()
}
