package controllers

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/althafariq/pemanduwisata-be/models"
	"github.com/gin-gonic/gin"
)

type CreateDestinationResponse struct {
	ID 		int 		`json:"id"`
	Name 	string 	`json:"name"`
	SuccessResponse
}

type DestinationRequest struct {
	ID 					int    `json:"id"` //don't need when create new destination
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

	if len(destination) == 0 {
		c.JSON(http.StatusOK, []string{}) //return empty array
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

func (api *API) GetDestination(c *gin.Context) {
	destinationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	destination, err := api.destinationModels.GetDestinationbyID(destinationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	avgRating, err := api.reviewModels.GetAverageRating(destinationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	totalReview, err := api.reviewModels.GetReviewCount(destinationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	destinationDetail := models.Destination{
		ID: 									destination[0].ID,
		Name: 								destination[0].Name,
		Location: 						destination[0].Location,
		Description: 					destination[0].Description,
		BudayaName: 					destination[0].BudayaName,
		BudayaDescription: 		destination[0].BudayaDescription,
		PhotoPath: 						destination[0].PhotoPath,
		AvgRating: 						avgRating,
		TotalReview: 					totalReview,
	}

	c.JSON(http.StatusOK, destinationDetail)
}

func (api *API) CreateDestination(c *gin.Context) {
	var req DestinationRequest
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

func (api *API) UploadImage(c *gin.Context) {
	var (
		destinationID int
		err           error
	)

	if destinationID, err = strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid destination ID"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	folderPath := "media/destination"
	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	files := form.File["images"]
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, file := range files {
		wg.Add(1)

		go func(file *multipart.FileHeader) {
			defer wg.Done()

			defer func() {
				if v := recover(); v != nil {
					log.Println(v)
					c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Internal Server Error"})
					return
				}
			}()

			uploadedFile, err := file.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
				return
			}

			defer uploadedFile.Close()

			unixTime := time.Now().UTC().UnixNano()
			fileName := fmt.Sprintf("%d-%d-%s", destinationID, unixTime, strings.ReplaceAll(file.Filename, " ", ""))
			fileLocation := filepath.Join(folderPath, fileName)
			targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)

			if err != nil {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
				return
			}

			defer targetFile.Close()

			if _, err := io.Copy(targetFile, uploadedFile); err != nil {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
				return
			}

			mu.Lock()
			if err := api.destinationModels.InsertDestinationImage(destinationID, fileLocation); err != nil {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
				return
			}
			mu.Unlock()
		}(file)
	}

	wg.Wait()

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Destination Images Uploaded",
	})

}

func (api *API) UpdateDestination(c *gin.Context) {
	var req DestinationRequest
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


