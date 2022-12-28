package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/althafariq/pemanduwisata-be/helper"
	"github.com/althafariq/pemanduwisata-be/models"
	"github.com/althafariq/pemanduwisata-be/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateReviewRequest struct {
	// DestinationID int    `json:"destination_id" binding:"required"` 
	Review string `json:"review" binding:"required"`
	Rating int    `json:"rating" binding:"required"`
}

type UpdateReviewRequest struct {
	ReviewID int    `json:"review_id" binding:"required"`
	Review string `json:"review" binding:"required"`
	Rating int    `json:"rating" binding:"required"`
}

func (api *API) GetAllReviews(c *gin.Context) {
	destinationID, err := strconv.Atoi(c.Param("destinationID"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	reviews, err := api.reviewModels.GetReviewbyDestinationID(destinationID)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	if len(reviews) == 0 {
		c.JSON(http.StatusOK, []string{}) //return empty array
		return
	}

	avgRating, err := api.reviewModels.GetAverageRating(destinationID)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	fmt.Println(reviews)

	reviewbyDestination := make(map[int]models.Review)
	for _, v := range reviews {
		if _, ok := reviewbyDestination[v.ID]; !ok {
			reviewbyDestination[v.ID] = models.Review{
				ID: v.ID,
				DestinationID: v.DestinationID,
				UserID: v.UserID,
				Firstname: v.Firstname,
				Lastname: v.Lastname,
				Profile_pic: v.Profile_pic,
				Rating: v.Rating,
				Review: v.Review,
				CreatedAt: v.CreatedAt,
			}
		}
	}

	res := make([]models.Review, 0, len(reviewbyDestination))
	for _, v := range reviewbyDestination {
		res = append(res, v)
	}

	c.JSON(http.StatusOK, gin.H{
		"avg_rating": avgRating,
		"reviews": res,
	})
}

func (api *API) CreateReview(c *gin.Context) {
	destinationID, err := strconv.Atoi(c.Param("destinationID"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": "error disini"},
		)
		return
	}

	var createReviewRequest CreateReviewRequest
	error := c.ShouldBindJSON(&createReviewRequest)
	if error != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			c.AbortWithStatusJSON(
				http.StatusBadRequest, 
				gin.H{"error": helper.GetErrorMessage(ve)},
			)
		} else {
			c.AbortWithStatusJSON(
				http.StatusBadRequest, 
				gin.H{"error": err.Error()},
			)
		}
		return
	}

	userID, err := api.getUserIdFromToken(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	isReviewOK := service.GetValidationInstance().Validate(createReviewRequest.Review)
	if !isReviewOK {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, 
			gin.H{"error": "Review contains bad words"}, 
		)
		return
	}

	reviewID, err := api.reviewModels.CreateReview(models.Review{
		DestinationID: destinationID,
		UserID: userID,
		Review: createReviewRequest.Review,
		Rating: createReviewRequest.Rating,
		//use time now
		CreatedAt: time.Now(),
	})
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "kenapa ya"},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Review added",
		"review_id": reviewID,
	})
}

func (api *API) DeleteReview(c *gin.Context) {
	reviewID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	userID, err := api.getUserIdFromToken(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	authorID, err := api.reviewModels.GetReviewbyUserID(reviewID)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	if authorID == 0 {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{"error": "No data with given id"},
		)
		return
	} else if authorID != userID {
		c.AbortWithStatusJSON(
			http.StatusForbidden,
			gin.H{"error": "You are not authorized to delete this"},
		)
		return
	}

	err = api.reviewModels.DeleteReview(reviewID)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{ "message": fmt.Sprintf("Review with id %d deleted", reviewID),
	})
}