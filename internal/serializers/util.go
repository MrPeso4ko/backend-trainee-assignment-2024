package serializers

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

func (b *BannerContent) Scan(value any) error {
	v, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(v, &b)
}
func (b BannerContent) Value() (driver.Value, error) {
	return json.Marshal(b)
}
