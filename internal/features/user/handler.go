package user

import (
	"time"

	"github.com/gin-gonic/gin"

	"boilerplate/internal/models"
	"boilerplate/internal/shared/constants"
	"boilerplate/internal/shared/dto"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type userResponse struct {
	ID        int64     `json:"user_id"`
	Name      string    `json:"name"`
	AuthID    string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type userCreate struct {
	Name string `json:"name" binding:"required"`
}

type userUpdate struct {
	Name string `json:"name" binding:"-"`
}

func (h *UserHandler) makeUserResponseFrom(user *models.User) *userResponse {
	return &userResponse{
		ID:        user.ID,
		Name:      user.Name,
		AuthID:    user.AuthID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	id := c.GetInt64(constants.UserIdKey)
	output, err := h.userService.GetUserById(id)

	if err != nil {
		dto.SendErrorResponse(c, err)
		return
	}

	dto.SendSuccessResponse(c, h.makeUserResponseFrom(output))
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input userCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		dto.SendErrorResponse(c, models.NewBadRequestError(err.Error()))
		return
	}

	user := models.User{
		Name:   input.Name,
		AuthID: c.GetString(constants.UserAuthIdKey),
	}
	output, err := h.userService.CreateUser(&user)

	if err != nil {
		dto.SendErrorResponse(c, err)
		return
	}

	dto.SendSuccessResponse(c, h.makeUserResponseFrom(output))
}

func (h *UserHandler) UpdateCurrentUser(c *gin.Context) {
	var input userUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		dto.SendErrorResponse(c, models.NewBadRequestError(err.Error()))
		return
	}

	user := models.User{
		ID:   c.GetInt64(constants.UserIdKey),
		Name: input.Name,
	}
	output, err := h.userService.UpdateUser(&user)

	if err != nil {
		dto.SendErrorResponse(c, err)
		return
	}

	dto.SendSuccessResponse(c, h.makeUserResponseFrom(output))
}

func (h *UserHandler) DeleteCurrentUser(c *gin.Context) {
	id := c.GetInt64(constants.UserIdKey)
	if err := h.userService.DeleteUserById(id); err != nil {
		dto.SendErrorResponse(c, err)
		return
	}

	dto.SendSuccessResponse(c, "User deleted.")
}
