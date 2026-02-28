package api

import (
	"fmt"
	"time"

	"github.com/Koubae/GoAnyBusiness/internal/app/core"
	"github.com/Koubae/GoAnyBusiness/internal/app/middleware"
	"github.com/Koubae/GoAnyBusiness/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigureRouter(router *gin.Engine, config *core.Config) error {
	allowOrigin := []string{"*"}
	allowALlOrigins := false
	if config.Env != core.Production {
		allowOrigin = nil
		allowALlOrigins = true
	}

	router.Use(
		cors.New(
			cors.Config{
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
				ExposeHeaders:    []string{"Content-Length"},
				MaxAge:           12 * time.Hour,
				AllowCredentials: false,
				AllowOrigins:     allowOrigin,
				AllowAllOrigins:  allowALlOrigins,
			},
		),
	)

	router.Use(middleware.RequestID())

	rps := utils.GetEnvFloat64("APP_RATE_LIMIT_RPS", 10)
	burst := utils.GetEnvInt("APP_RATE_LIMIT_BURST", 20)
	rl := middleware.NewRateLimiter(rps, burst)
	router.Use(rl.Limit())

	err := router.SetTrustedProxies(config.TrustedProxies)
	if err != nil {
		return fmt.Errorf("Error setting trusted proxies, error: %s", err.Error())
	}

	index := router.Group("/")
	indexController := &IndexController{
		config: config,
	}
	{
		index.GET("/", indexController.Index)
		index.GET("/ping", indexController.Ping)
		index.GET("/alive", indexController.Alive)
		index.GET("/ready", indexController.Ready)
	}

	healthController := &HealthController{
		config: config,
	}
	index.GET("/health", healthController.Health)

	return nil
}
