package controllers

import (
	"net/http"
	"strconv"

	"github.com/althafariq/pemanduwisata-be/models"
	"github.com/gin-gonic/gin"
)

type CreateDestinationRequest struct {
	Name        string `json:"name" binding:"required"`
	Location    string `json:"location" binding:"required"`
	Description string `json:"description" binding:"required"`
	BudayaName  string `json:"budaya_name"`
	BudayaDescription string `json:"budaya_description"`
	PhotoPath   string `json:"photo_path"`
}

type CreateDestinationResponse struct {
	ID 		int 		`json:"id"`
	Name 	string 	`json:"name"`
	SuccessResponse
}

type UpdateDestinationRequest struct {
	ID 					int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	Location    string `json:"location" binding:"required"`
	Description string `json:"description" binding:"required"`
	BudayaName  string `json:"budaya_name"`
	BudayaDescription string `json:"budaya_description"`
	PhotoPath   string `json:"photo_path"`
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
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	destinationDetail := make(map[int]models.Destination)
	for _, v := range destination {
		if _, ok := destinationDetail[v.ID]; !ok {
			avgRating, err := api.reviewModels.GetAverageRating(v.ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
				return
			}

			totalReview, err := api.reviewModels.GetReviewCount(v.ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
				return
			}

				destinationDetail[v.ID] = models.Destination{
					ID: 									v.ID,
					Name: 								v.Name,
					Location: 						v.Location,
					Description: 					v.Description,
					BudayaName: 					v.BudayaName,
					BudayaDescription: 		v.BudayaDescription,
					PhotoPath: 						v.PhotoPath,
					AvgRating: 						avgRating,
					TotalReview: 					totalReview,
				}
			} 
		}

	res := make([]models.Destination, 0, len(destinationDetail)) // pre-allocate the slice with the right capacity
	for _, v := range destinationDetail {
		res = append(res, v)
	}

	c.JSON(http.StatusOK, res)
}

// func (api *API) GetDestination(c *gin.Context) {
// 	// load all information about destination such as reviews, average rating, total review, and reviews
// 	destinationID, err := strconv.Atoi(c.Query("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
// 		return
// 	}

// 	destination, err := api.destinationModels.GetDestinationbyID(destinationID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
// 		return
// 	}

// 	avgRating, err := api.reviewModels.GetAverageRating(destinationID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
// 		return
// 	}

// 	totalReview, err := api.reviewModels.GetReviewCount(destinationID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
// 		return
// 	}

// 	reviewComment, err := api.reviewModels.GetReviewbyDestinationID(destinationID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
// 		return
// 	}

// 	destinationDetail := models.Destination{
// 		ID: 									destination.ID,
// 		Name: 								destination.Name,
// 		Location: 						destination.Location,
// 		Description: 					destination.Description,
// 		BudayaName: 					destination.BudayaName,
// 		BudayaDescription: 		destination.BudayaDescription,
// 		PhotoPath: 						destination.PhotoPath,
// 		AvgRating: 						avgRating,
// 		TotalReview: 					totalReview,
// 		Review: 							reviewComment[0],
// 	}

// 	c.JSON(http.StatusOK, destinationDetail)
// }

// create destination
func (api *API) CreateDestination(c *gin.Context) {
	var req CreateDestinationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	destinationID, err := api.destinationModels.CreateDestination(models.Destination{
		Name:        				req.Name,
		Location:    				req.Location,
		Description: 				req.Description,
		BudayaName:  				req.BudayaName,
		BudayaDescription: 	req.BudayaDescription,
		PhotoPath:   				req.PhotoPath,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CreateDestinationResponse{
		ID: destinationID,
		Name: req.Name,
		SuccessResponse: SuccessResponse{
			Message: "Destination created successfully",
		},
	})
}

// update destination
func (api *API) UpdateDestination(c *gin.Context) {
	var req UpdateDestinationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	err := api.destinationModels.UpdateDestination(
		req.ID,
		req.Name,
		req.Location,
		req.Description,
		req.BudayaName,
		req.BudayaDescription,
		req.PhotoPath,
	)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}	

	c.JSON(http.StatusOK, CreateDestinationResponse{
		ID: req.ID,
		Name: req.Name,
		SuccessResponse: SuccessResponse{
			Message: "Destination updated successfully",
		},
	})
}

// delete destination
func (api *API) DeleteDestination(c *gin.Context) {
	destinationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	if err := api.destinationModels.DeleteDestination(destinationID); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Destination deleted successfully"})
}


