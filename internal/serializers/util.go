package serializers

import (
	"encoding/json"
	"fmt"
)

type dbSlice []int

func (s *dbSlice) Scan(src any) error {
	fmt.Println("HERE")
	var srcBytes []byte
	switch v := src.(type) {
	case []byte:
		srcBytes = v
	case string:
		srcBytes = []byte(v)
	}
	return json.Unmarshal(srcBytes, s)
}

func (b *BannerContent) Scan(src any) error {
	fmt.Println("HERE")
	var srcBytes []byte
	switch v := src.(type) {
	case []byte:
		srcBytes = v
	case string:
		srcBytes = []byte(v)
	}
	return json.Unmarshal(srcBytes, b)
}

type bannerContentString struct {
	string
}

func (s *bannerContentString) UnmarshalJSON(b []byte) error {
	s.string = string(b)
	return nil
}
