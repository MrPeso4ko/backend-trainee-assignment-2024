package api

import (
	"backend-trainee-assignment-2024/m/internal/config"
	"backend-trainee-assignment-2024/m/pkg/logging"
	"backend-trainee-assignment-2024/m/pkg/postgres"
	"github.com/gin-gonic/gin"
)

var logger = logging.GetLogger()

func StartAPIServer() {
	logger.Info("Starting API server...")

	db, err := postgres.Connect(config.GetPostgresDSN())
	if err != nil {
		logger.Errorf("error connecting to postgres: %+v", err)
		return
	}

	r := gin.Default()
	err = r.SetTrustedProxies(nil)
	if err != nil {
		logger.Errorf("error setting trusted proxies: %+v", err)
		return
	}

	r.GET("/user_banner", func(c *gin.Context) {
		getUserBanner(db, c)
	})

	adminGroup := r.Group("")
	adminGroup.GET("/banner", func(c *gin.Context) {
		getBanners(db, c)
	})
	adminGroup.POST("/banner", func(c *gin.Context) {
		createBanner(db, c)
	})

	err = r.Run(config.GetAPIAddress())
	if err != nil {
		logger.Errorf("Error starting server: %+v", err)
		return
	}
}
