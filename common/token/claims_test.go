package token

import (
	"testing"
	"time"
)

func TestClaims_Valid(t *testing.T) {
	expired := 2 * time.Hour
	t.Run("过期", func(t *testing.T) {
		claims := Claims{
			ExpiresAt: time.Now().Unix() + int64(expired.Seconds()) - int64((3 * time.Hour).Seconds()),
		}
		if err := claims.Valid(); err.Error() == "token is expired" {
			t.Log("已过期")
		}
	})

	t.Run("在时效内", func(t *testing.T) {
		claims := Claims{
			ExpiresAt: time.Now().Unix() + int64(expired.Seconds()) - int64((1 * time.Hour).Seconds()),
		}
		if err := claims.Valid(); err == nil {
			t.Log("在时效内")
		}
	})
}
