package jwtauth

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type (
	UserInfo struct {
		UserID uint64
	}

	UserInfoClaims struct {
		UserInfo
		jwt.RegisteredClaims
	}
)

// CreateToken
// @param secret
// @param userinfo
// @date 2023-05-06 17:49:17
func CreateToken(secret string, userinfo UserInfo) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserInfoClaims{
		UserInfo: userinfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
	})
	return token.SignedString([]byte(secret))
}
