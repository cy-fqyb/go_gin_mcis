package controller

import "github.com/gin-gonic/gin"

func init() {
	RegisterPrivateRoutes(
		func(r *gin.RouterGroup) {
			r.GET("/tzfm", tzfmHandler)
		},
	)
}

func tzfmHandler(c *gin.Context) {

}
