package middleware

import (
	"boilerplate/internal/features/auth"
	"boilerplate/internal/features/user"
	"boilerplate/internal/models"
	"boilerplate/internal/shared/constants"
	"boilerplate/internal/shared/dto"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authService auth.AuthService
	userService user.UserService
}

func NewAuthMiddleware(
	authService auth.AuthService,
	userService user.UserService,
) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
		userService: userService,
	}
}

// TODO: Add require verification implementation
func (am *AuthMiddleware) AuthMiddleware(c *gin.Context) {
	authorizationToken := c.GetHeader("Authorization")
	authToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

	// No token provided
	if authToken == "" {
		c.Next()
		return
	}

	// Validate token
	authId, validateErr := am.authService.Validate(authToken)
	if validateErr != nil {
		dto.SendErrorResponse(c, validateErr)
		return
	}

	c.Set(constants.UserAuthIdKey, authId)

	userId, getUserIdErr := am.userService.GetUserIdByAuthId(authId)
	if getUserIdErr != nil {
		c.Next()
		return
	}

	c.Set(constants.UserIdKey, userId)
	c.Next()
}

func (am *AuthMiddleware) UserRequiredMiddleware(c *gin.Context) {
	if authId, exists := c.Get(constants.UserAuthIdKey); authId == "" || !exists {
		dto.SendErrorResponse(c, models.NewUnauthorizedError(false))
		return
	}

	// For the case when firebase user created but app user not yet created
	// Can be caused by error or front end's flow
	if userId, exists := c.Get(constants.UserIdKey); userId == "" || !exists {
		dto.SendErrorResponse(c, models.NewBadRequestError("Finish setting up your account first."))
		return
	}

	c.Next()
}
