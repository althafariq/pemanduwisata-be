package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) GetAllTelpDarurat(c *gin.Context) {
	telpDarurat, err := api.telpDaruratModels.GetAllTelpDarurat()
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, telpDarurat)
}
