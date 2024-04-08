package storage

import (
	"backend-trainee-assignment-2024/m/internal/serializers"
	"github.com/jmoiron/sqlx"
)

func GetBanner(db *sqlx.DB, featureId, tagId int) (*serializers.Banner, error) {
	sql := `SELECT title, text, url
FROM banners_manager.banners
WHERE feature_id = $1
  AND $2 = ANY (tag_ids) LIMIT 1`
	var res []serializers.Banner
	err := db.Select(&res, sql, featureId, tagId)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	return &res[0], nil
}
