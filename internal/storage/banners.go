package storage

import (
	"backend-trainee-assignment-2024/m/internal/serializers"
	"github.com/jmoiron/sqlx"
)

func GetUserBanner(db *sqlx.DB, request serializers.UserBannerRequest) (*serializers.BannerContent, error) {
	sql := `SELECT content
FROM banners_manager.banners
WHERE feature_id = $1
  AND $2 = ANY (tag_ids) LIMIT 1`
	var res []*serializers.BannerContent
	err := db.Select(&res, sql, request.FeatureId, request.TagId)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	return res[0], nil
}

func GetBanners(db *sqlx.DB, request serializers.BannerRequest) ([]*serializers.BannerGet, error) {
	sql := `SELECT id, tag_ids as tag_ids, feature_id, content, is_active, created_at, updated_at
FROM banners_manager.banners
WHERE ($1::INT IS NULL OR feature_id = $1)
  AND ($2::INT IS NULL OR $2 = ANY (tag_ids)) ORDER BY id LIMIT COALESCE($3, 50) OFFSET COALESCE($4, 0)`
	var res []*serializers.BannerGet
	err := db.Select(&res, sql, request.FeatureId, request.TagId, request.Limit, request.Offset)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return []*serializers.BannerGet{}, nil
	}
	return res, nil
}

func CreateBanner(db *sqlx.DB, banner serializers.BannerCreate) (int, error) {
	sql := `INSERT INTO banners_manager.banners(feature_id, tag_ids, is_active, content)
VALUES (:feature_id, :tag_ids, :is_active, :content)
RETURNING id;`
	var id int
	rows, err := db.NamedQuery(sql, banner)
	if err != nil {
		return -1, err
	}
	rows.Next()
	err = rows.Scan(&id)
	return id, err
}
