package utils

import "time"

const (
	LayoutDateTime = "2006-01-02 15:04:05"
)

func VerifyExpiresAt(t int64) bool {
	return time.Now().Unix() <= t
}

func GetDateTime(str string) (time.Time, error) {
	return time.Parse(LayoutDateTime, str)
}
