package helper

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TimeToMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
func DataTimeToTime(dt primitive.DateTime) time.Time {
	return time.Unix(0, int64(dt)*int64(time.Millisecond))
}
