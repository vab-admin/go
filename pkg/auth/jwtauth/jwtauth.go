package jwtauth

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type (
	UserInfo struct {
		UserID uint64
	}

	UserInfoClaims struct {
		UserInfo
		jwt.StandardClaims
	}
)

// CreateToken
// @param secret
// @param userinfo
// @date 2023-05-06 17:49:17
func CreateToken(secret string, userinfo UserInfo) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserInfoClaims{
		UserInfo: userinfo,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	})
	return token.SignedString([]byte(secret))
}
