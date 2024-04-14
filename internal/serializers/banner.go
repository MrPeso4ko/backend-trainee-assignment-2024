package serializers

import (
	"github.com/lib/pq"
	"time"
)

type UserBannerRequest struct {
	FeatureId       int  `form:"feature_id" binding:"required"`
	TagId           int  `form:"tag_id" binding:"required"`
	UseLastRevision bool `form:"use_last_revision"`
}

type BannerRequest struct {
	FeatureId *int `form:"feature_id"`
	TagId     *int `form:"tag_id"`
	Limit     *int `form:"limit"`
	Offset    *int `form:"offset"`
}

type BannerContent map[string]any
type BannerBase struct {
	TagIds    pq.Int32Array `json:"tag_ids" db:"tag_ids"`
	FeatureID int           `json:"feature_id" db:"feature_id"`
	Content   BannerContent `json:"content" db:"content"`
	IsActive  bool          `json:"is_active" db:"is_active"`
}
type BannerGet struct {
	Id        int       `json:"banner_id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	BannerBase
}

type BannerUpdate struct {
	BannerBase
}

type BannerCreate struct {
	TagIds    pq.Int32Array `json:"tag_ids" db:"tag_ids" binding:"required"`
	FeatureID int           `json:"feature_id" db:"feature_id" binding:"required"`
	Content   BannerContent `json:"content" db:"content" binding:"required"`
	IsActive  bool          `json:"is_active" db:"is_active" binding:"required"`
}
