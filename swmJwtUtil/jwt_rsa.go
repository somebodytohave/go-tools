package swmJwtUtil

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

type CustomClaims struct {
	Data interface{} `json:"data"`
	jwt.StandardClaims
}

// GenRSA256TokenWithDay 默认一天保存时间
// filename: 密钥的名称
func GenRSA256TokenWithDay(claims CustomClaims, filename string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims.ExpiresAt = expireTime.Unix()
	// 私钥
	key, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		return "", err
	}
	return GenRSA256TokenWithPriKey(claims, privateKey)
}

// GenRSA256TokenWithFileName 默认一天保存时间
// filename: 密钥的名称
func GenRSA256TokenWithFileName(claims CustomClaims, filename string) (string, error) {
	// 私钥
	key, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		return "", err
	}
	return GenRSA256TokenWithPriKey(claims, privateKey)
}

// GenRSA256TokenWithPriKey 生成 加密方式 token
func GenRSA256TokenWithPriKey(claims CustomClaims, privateKey *rsa.PrivateKey) (string, error) {
	// claims.Issuer = "sun-wenming@secure.istio.io"
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	//header := map[string]interface{}{
	//	"alg": "RS256", "typ": "JWT",
	//}
	//tokenClaims.Header = header
	// 获取完整签名之后的 token
	return tokenClaims.SignedString(privateKey)
}

// ParseRAS256TokenFileName
// filename: 密钥的名称
func ParseRAS256TokenFileName(token string, filename string) (*CustomClaims, error) {
	// 私钥
	key, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(key)
	if err != nil {
		return nil, err
	}
	return ParseRAS256TokenPubKey(token, publicKey)
}

// ParseRAS256TokenFileName 解析token
func ParseRAS256TokenPubKey(token string, publicKey *rsa.PublicKey) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if !tokenClaims.Valid {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil, errors.New("timing is everything")
			} else {
				return nil, errors.New("couldn't handle this token:" + err.Error())
			}
		} else {
			return nil, errors.New("couldn't handle this token:" + err.Error())
		}
	}

	if tokenClaims == nil {
		return nil, errors.New("token解析失败")
	}

	if claims, ok := tokenClaims.Claims.(*CustomClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("token Claims is not CustomClaims")
	}
}
