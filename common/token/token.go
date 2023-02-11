package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"Reward/common"
	"Reward/log"
)

var jwtKey string

type PayLoad struct {
	UserID  int           `json:"user_id"`
	Expired time.Duration `json:"expired"`
}

type ParsedPayLoad struct {
	UserID    int   `json:"user_id"`
	ExpiresAt int64 `json:"expires_at"` // 过期时间（时间戳，10位）
}

func getJwtKey() string {
	if jwtKey == common.StringEmpry {
		jwtKey = viper.GetString("jwt_secret")
	}
	return jwtKey
}

// 根据用户的用户名和密码产生token
func GenerateToken(payload *PayLoad) (string, error) {
	//设置token有效时间
	claims := &Claims{
		UserID:    payload.UserID,
		ExpiresAt: time.Now().Unix() + int64(payload.Expired.Seconds()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	return token.SignedString([]byte(getJwtKey()))
}

// 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func ParseToken(tokenStr string) (*ParsedPayLoad, error) {
	claims := &Claims{}
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(getJwtKey()), nil
	})

	if err != nil {
		log.Error("token parsing failed because of an internal error", zap.String("cause", err.Error()))
		return nil, err
	}

	// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
	// 要传入指针，项目中结构体都是用指针传递，节省空间。
	if !token.Valid {
		log.Error("Token parsing failed; the token is invalid.")
		return nil, errors.New("invalid token")
	}

	t := &ParsedPayLoad{
		UserID:    claims.UserID,
		ExpiresAt: claims.ExpiresAt,
	}

	return t, nil
}
