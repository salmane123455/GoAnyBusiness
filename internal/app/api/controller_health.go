package api

import (
	"net/http"
	"time"

	"github.com/Koubae/GoAnyBusiness/internal/app/core"
	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

type HealthController struct {
	config *core.Config
}

func (h *HealthController) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":      "healthy",
		"uptime":      time.Since(startTime).String(),
		"version":     h.config.AppVersion,
		"environment": h.config.Env,
	})
}
