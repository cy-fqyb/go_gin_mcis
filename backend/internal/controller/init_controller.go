package controller

import "github.com/gin-gonic/gin"

type InitControllerFunc func(r *gin.RouterGroup)

var (
	publicRoutes  []InitControllerFunc
	privateRoutes []InitControllerFunc
)

func InitController(r *gin.RouterGroup) {
	publicPrefix := r.Group("/public")
	privatePrefix := r.Group("/")

	for _, route := range publicRoutes {
		route(publicPrefix)
	}

	for _, route := range privateRoutes {
		route(privatePrefix)
	}
}

func RegisterPublicRoutes(f InitControllerFunc) {
	publicRoutes = append(publicRoutes, f)
}

func RegisterPrivateRoutes(f InitControllerFunc) {
	privateRoutes = append(privateRoutes, f)
}
