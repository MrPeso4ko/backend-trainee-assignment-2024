package api

import (
	"backend-trainee-assignment-2024/m/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

func getUserBanner(db *sqlx.DB, context *gin.Context) {
	logger.Info("getting segments...")

	featureId, err := strconv.Atoi(context.Query("feature_id"))
	if err != nil {
		logger.Errorf("Can't convert feature_id to int: %+v", err)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid feature_id"})
		return
	}

	tagId, err := strconv.Atoi(context.Query("tag_id"))
	if err != nil {
		logger.Errorf("Can't convert tag_id to int: %+v", err)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid tag_id"})
		return
	}
	banner, err := storage.GetBanner(db, featureId, tagId)
	if err != nil {
		logger.Errorf("Can't get banner for %d/%d : %+v", featureId, tagId, err)
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": true, "message": "internal error"})
		return
	}
	if banner == nil {
		context.AbortWithStatus(http.StatusNotFound)
		return
	}

	context.JSON(http.StatusOK, banner)
	logger.Info("success")
}
