package dto

import (
	"boilerplate/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func SendSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Data:  data,
		Error: nil,
	})
}

func SendErrorResponse(c *gin.Context, err *models.AppError) {
	c.AbortWithStatusJSON(
		err.Status(),
		Response{
			Data:  nil,
			Error: err.Error(),
		},
	)
}
