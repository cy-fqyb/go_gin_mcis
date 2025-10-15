package controller

import (
	"encoding/json"
	"go_gin_mcis/internal/dto"
	"go_gin_mcis/internal/service"
	"go_gin_mcis/pkg/logger"
	"go_gin_mcis/pkg/result"
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
		result.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	queryJson, _ := json.Marshal(query)
	logger.Info("Received query: ", string(queryJson))

	records, err := service.GetUploadRecordByType(*query.UploadType, query)
	if err != nil {
		result.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	result.Success(c, records)
}
