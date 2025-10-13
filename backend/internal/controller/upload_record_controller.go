package controller

import (
	"go_gin_mcis/internal/dto"
	"go_gin_mcis/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	RegisterPrivateRoutes(func(r *gin.RouterGroup) {
		r.GET("/upload_records", GetUploadRecords)
	})
}

func GetUploadRecords(c *gin.Context) {
	var query dto.UploadRecordQuery

	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	records, err := service.GetUploadRecord(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": records,
	})
}
