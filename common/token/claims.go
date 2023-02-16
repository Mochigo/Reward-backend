package token

import (
	"errors"

	"Reward/common/utils"
)

//Claim是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	UserID    int   `json:"user_id"`
	CollegeId int   `json:"college_id"`
	ExpiresAt int64 `json:"expires_at"` // 过期时间（时间戳，10位）
}

func (c *Claims) Valid() error {
	if !utils.VerifyExpiresAt(c.ExpiresAt) {
		return errors.New("token is expired")
	}
	return nil
}
