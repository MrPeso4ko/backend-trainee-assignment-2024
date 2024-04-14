package api

import (
	"backend-trainee-assignment-2024/m/internal/serializers"
	"backend-trainee-assignment-2024/m/internal/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func getUserBanner(db *sqlx.DB, context *gin.Context) {
	logger.Info("Getting user banner")

	req := serializers.UserBannerRequest{}
	err := context.ShouldBindQuery(&req)
	if err != nil {
		logger.Errorf("Couldn't bind query: %+v", err)
		context.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("Wrong argument: %s", err.Error())})
		return
	}
	logger.Debugf("Got request: %+v", req)
	banner, err := storage.GetUserBanner(db, req)
	if err != nil {
		logger.Errorf("Can't get banner for %d/%d : %+v", req.FeatureId, req.TagId, err)
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	if banner == nil {
		context.AbortWithStatus(http.StatusNotFound)
		return
	}

	context.JSON(http.StatusOK, banner)
	logger.Info("success")
}

func getBanners(db *sqlx.DB, context *gin.Context) {
	logger.Info("Getting user banner")

	req := serializers.BannerRequest{}
	err := context.ShouldBindQuery(&req)
	if err != nil {
		logger.Errorf("Couldn't bind query: %+v", err)
		context.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("Wrong argument: %s", err.Error())})
		return
	}
	logger.Debugf("Got request: %+v", req)
	banners, err := storage.GetBanners(db, req)
	if err != nil {
		logger.Errorf("Can't get banners for %v/%v : %+v", req.FeatureId, req.TagId, err)
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	context.JSON(http.StatusOK, banners)
	logger.Info("success")
}

func createBanner(db *sqlx.DB, context *gin.Context) {
	logger.Info("creating banner")

	banner := serializers.BannerCreate{}
	err := context.ShouldBindJSON(&banner)
	if err != nil {
		logger.Errorf("Couldn't bind JSON: %+v", err)
		context.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("Wrong argument: %s", err.Error())})
		return
	}

	bannerId, err := storage.CreateBanner(db, banner)
	if err != nil {
		logger.Errorf("Can't create banner %+v: %+v", banner, err)
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"banner_id": bannerId})
}
