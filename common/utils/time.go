package utils

import "time"

func VerifyExpiresAt(t int64) bool {
	return time.Now().Unix() <= t
}
