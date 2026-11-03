package api

import (
	"log"

	"github.com/Radikahn/gluster/api/routes/devices"
	"github.com/Radikahn/gluster/api/routes/system"
	"github.com/gin-gonic/gin"
)

func StartOrchestratorAPI(logger log.Logger) {
	logger.Print("[ORCHESTRATOR]: Starting Orchestrator API...")
	router := gin.Default()

	// base route!
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from gluster",
		})
	})

	api := router.Group("/api")
	system.RegisterRoutes(api.Group("/system"))
	devices.RegisterRoutes(api.Group("/devices"))

	router.Run()
	logger.Print("[ORCHESTRATOR]: Api Started!")
}
