package utils

import "time"

var (
	Tokyo *time.Location
)

func init() {
	tz, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	Tokyo = tz
}
