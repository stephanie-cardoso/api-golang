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

// @BasePath /

// @Summary Get opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} GetOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func GetOpening(c *gin.Context) {
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

	sendSuccess(c, "get-opening", opening)
}

// @BasePath /

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpening(c *gin.Context) {
	newOpening := CreateOpeningRequest{}
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
	sendSuccess(c, "create-opening", opening)
}

// @BasePath /

// @Summary Delete opening
// @Description Delete a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
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
	sendSuccess(c, "delete-opening", opening)
}

// @BasePath /

// @Summary Update opening
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening Identification"
// @Param opening body UpdateOpeningRequest true "Opening data to Update"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpening(c *gin.Context) {
	receivedOpening := UpdateOpeningRequest{}

	c.BindJSON(&receivedOpening)

	if err := receivedOpening.Validate(); err != nil {
		log.Error().Err(err).Msgf("[handler] validation error to update opening")
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	updatingOpening := schemas.Opening{}
	if err := db.First(&updatingOpening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "opening not found")
		return
	}

	updatedOpening := updateOpening(receivedOpening, updatingOpening)
	fmt.Println("PUT OPENING:", updatedOpening)
	if err := db.Save(&updatedOpening).Error; err != nil {
		log.Error().Err(err).Msgf("[handler] error updating opening")
		sendError(c, http.StatusInternalServerError, "error updating opening")
		return
	}

	sendSuccess(c, "update-opening", updatingOpening)
}

// @BasePath /

// @Summary Get openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func GetOpenings(c *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprint("failed to list openings"))
		return
	}

	sendSuccess(c, "get-openings", openings)
}

func updateOpening(receivedOpening UpdateOpeningRequest, opening schemas.Opening) schemas.Opening {
	if receivedOpening.Role != "" {
		opening.Role = receivedOpening.Role
	}

	if receivedOpening.Company != "" {
		opening.Company = receivedOpening.Company
	}

	if receivedOpening.Location != "" {
		opening.Location = receivedOpening.Location
	}

	if receivedOpening.Remote != nil {
		opening.Remote = *receivedOpening.Remote
	}

	if receivedOpening.Link != "" {
		opening.Link = receivedOpening.Link
	}

	if receivedOpening.Salary > 0 {
		opening.Salary = receivedOpening.Salary
	}

	return opening
}
