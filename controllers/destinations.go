package controllers

import (
	"net/http"

	"github.com/althafariq/pemanduwisata-be/models"
	"github.com/gin-gonic/gin"
)

type CreateDestinationRequest struct {
	Name        string `json:"name" binding:"required"`
	Location    string `json:"location" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreateDestinationResponse struct {
	ID int `json:"id"`
	SuccessResponse
}

type DestinationResponse struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Location      string  `json:"location"`
	Description   string  `json:"description"`
	BudayaID      int     `json:"budaya_id"`
	PhotoID       int     `json:"photo_id"`
	ReviewCount   int     `json:"review_count"`
	AverageRating float64 `json:"average_rating"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"error"`
}

func (api *API) GetAllDestination(c *gin.Context) {
	destination, err := api.destinationModels.GetAllDestinations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, 
		gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, destination)
}

func (api *API) CreateDestination(c *gin.Context) {
	var (
		req = CreateDestinationRequest{}
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request body"})
		return
	}

	destination := models.Destination{
		Name:        req.Name,
		Location:    req.Location,
		Description: req.Description,
	}

	if err := api.destinationModels.CreateDestination(&destination).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, CreateDestinationResponse{
		ID:              destination.ID,
		SuccessResponse: SuccessResponse{Message: "Destination created successfully"},
	})
}

