package main

import (
	"github.com/Radikahn/gluster/api/routes/devices"
	"github.com/Radikahn/gluster/api/routes/system"
	"github.com/gin-gonic/gin"
)

func main() {
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
}
