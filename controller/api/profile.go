package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"boilerplate/constants"
	"boilerplate/models"
	"boilerplate/services"
)

const profileNotFoundMessage = "profile not found"
const profileUpdateErrorMessage = "error when updating profile"

func CreateProfile(c *gin.Context) {
	db := services.Database

	uuid := c.GetString(constants.UserIDKey)

	var input models.ProfileCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
		return
	}

	profile := models.Profile{
		ID:        uuid,
		Name:      input.Name,
		BirthDate: input.BirthDate,
	}
	db.Create(&profile)

	c.JSON(http.StatusOK, models.Response{Data: profile})
}

func UpdateProfile(c *gin.Context) {
	db := services.Database

	uuid := c.GetString(constants.UserIDKey)

	var input models.ProfileUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
		return
	}

	var profile models.Profile
	if result := db.First(&profile, "id = ?", uuid); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: profileNotFoundMessage})
		return
	}

	if result := db.Model(&profile).Updates(models.Profile{
		Name:      input.Name,
		BirthDate: input.BirthDate,
	}); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: profileUpdateErrorMessage})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: profile})
}

func GetMyProfile(c *gin.Context) {
	db := services.Database

	uuid := c.Query(constants.UserIDKey)
	if uuid == "" {
		uuid = c.GetString(constants.UserIDKey)
	}

	var profile models.Profile
	if result := db.First(&profile, "id = ?", uuid); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: profileNotFoundMessage})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: profile})
}

func DeleteProfile(c *gin.Context) {
	db := services.Database

	uuid := c.GetString(constants.UserIDKey)

	if result := db.Delete(&models.Profile{}, "id = ?", uuid); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: profileNotFoundMessage})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: "Profile deleted"})
}
