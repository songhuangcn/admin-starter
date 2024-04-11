package auth

import (
	"errors"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/golang-jwt/jwt"
)

func JwtEncode(uid uint, jwtExpHours int, secret string) (string, error) {
	exp := time.Now().Add(time.Duration(jwtExpHours) * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": exp,
	})

	return token.SignedString([]byte(secret))
}

func JwtDecode(tokenString string, secret string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证加密方式是否一致
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("非法 Token")
		}

		return []byte(secret), nil
	})
	// 验证有效期等限制条件是否符合要求
	if err != nil {
		log.Debugf("Parse 失败：%#v", err.Error())
		return 0, errors.New("无效 Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Debugf("转化对象失败：%#v", token.Claims)
		return 0, errors.New("无效 Token")
	}

	log.Debugf("claims：%#v", claims)
	uid, ok := claims["uid"].(float64) // JWT 标准规范中规定，JSON 中的数字默认被解析为浮点数
	if !ok {
		log.Debugf("转化 uint 失败：%#v", claims["uid"])
		return 0, errors.New("无效 Token")
	}

	return uint(uid), nil
}
