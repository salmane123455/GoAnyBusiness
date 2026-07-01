package api

import (
	"fmt"
	"net/http"

	"github.com/Koubae/GoAnyBusiness/internal/app/core"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	config *core.Config
}

func (controller *IndexController) Index(c *gin.Context) {
	Success(c, fmt.Sprintf("Welcome to %s V%s", controller.config.AppName, controller.config.AppVersion), nil)
}

func (controller *IndexController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (controller *IndexController) Alive(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (controller *IndexController) Ready(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
