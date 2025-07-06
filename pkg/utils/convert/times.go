package convert

import "time"

func FromTimestampToString(t int64) string {
	return time.Unix(t, 0).UTC().Format(time.RFC3339)
}

func FromTimestampToTimeTime(t int64) time.Time {
	return time.Unix(t, 0).UTC()
}

func FromTimeTimeToTimestamp(t time.Time) int64 {
	return t.Unix()
}
