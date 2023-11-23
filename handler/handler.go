package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/stephanie-cardoso/api-golang/config"
	"github.com/stephanie-cardoso/api-golang/schemas"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeHandler() {
	db = config.GetSQLite()
}

func GetOpening(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get opening",
	})
}

func NewOpening(c *gin.Context) {
	newOpening := NewOpeningRequest{}
	c.BindJSON(&newOpening)

	if err := newOpening.Validate(); err != nil {
		log.Error().Err(err).Msgf("[handler] validation error")
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     newOpening.Role,
		Company:  newOpening.Company,
		Location: newOpening.Location,
		Remote:   *newOpening.Remote,
		Link:     newOpening.Link,
		Salary:   newOpening.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		log.Error().Err(err).Msgf("[handler] failed to create opening")
		sendError(c, http.StatusInternalServerError, "error creating opening in database")
		return
	}
	sendSucess(c, "create-opening", opening)
}

func DeleteOpening(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("opening with id %s not found in database", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("failed to delete opening with id %s", id))
		return
	}
	sendSucess(c, "delete-opening", opening)
}

func UpdateOpening(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "put opening",
	})
}

func GetOpenings(c *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprint("failed to list openings"))
		return
	}

	sendSucess(c, "get-openings", openings)
}
