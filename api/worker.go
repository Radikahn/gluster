package api

import (
	"log"

	"github.com/Radikahn/gluster/api/routes/system"
	"github.com/gin-gonic/gin"
)

type CommandInstance struct {
	command string `json:"command"`
}

func StartWorkerAPI(logger log.Logger) {
	logger.Print("[WORKER]: Starting Orchestrator API...")
	router := gin.Default()

	// base route!
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from gluster",
		})
	})

	api := router.Group("/api")
	system.RegisterRoutes(api.Group("/system"))

	router.Run()
	logger.Print("[WORKER]: Api Started!")
}
