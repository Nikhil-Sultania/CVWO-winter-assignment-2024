package router

import (
	"github.com/CVWO/sample-go-app/internal/routes"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	routes.RegisterRoutes(r)
	return r
}
