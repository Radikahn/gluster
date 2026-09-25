package system

import (
	"net/http"

	"github.com/Radikahn/gluster/api/types"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/health", Health)
}

// Health endpoint, returns status message!
//
// {status: message string}
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, types.HealthResponse{Status: "Healthy"})
}
