package api

import (
	"log"
	"net/http"

	"github.com/Radikahn/gluster/api"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/command/execute", executeCommand)
}

func executeCommand(c *gin.Context) {
	var command api.CommandInstance

	if err := c.BindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
		return
	}

	//TODO: process command here
	log.Print("Recieved the following command: ", command.Command)

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Command executed succesfully",
		"data":    command.Command,
	})
}
