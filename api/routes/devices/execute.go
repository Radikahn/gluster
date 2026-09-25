package devices

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/execute/:command", executeCommand)
}

func executeCommand(c *gin.Context) {
	// TODO: on failure, respond with http.StatusForbidden {"status": "failed"} and return early

	//TODO: do something with the command argument that comes in
	command := c.Param("command")

	slog.Info("[API]: running command:", "command", command)

	c.JSON(http.StatusOK, gin.H{
		"status": "complete",
	})

}
