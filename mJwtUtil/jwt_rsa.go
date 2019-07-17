package mJwtUtil

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

type CustomClaims struct {
	Data []byte `json:"data"`
	jwt.StandardClaims
}

// GenRSA256TokenWithFileName
// filename: 密钥的名称
func GenRSA256TokenWithFileName(claims CustomClaims, priKeyName, pubKeyName string) (string, error) {
	// 私钥
	priKey, err := ioutil.ReadFile(priKeyName)
	if err != nil {
		return "", err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(priKey)
	if err != nil {
		return "", err
	}
	// 公钥
	pubKey, err := ioutil.ReadFile(pubKeyName)
	if err != nil {
		return "", err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)
	if err != nil {
		return "", err
	}
	return GenRSA256Token(claims, privateKey, publicKey)
}

// GenRSA256Token 生成 加密方式 token
func GenRSA256Token(claims CustomClaims, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (string, error) {
	bytes, err := EncryptWithPublicKey(claims.Data, publicKey)
	if err != nil {
		return "", err
	}

	claims.Data = bytes
	// 默认 一天过期
	if claims.ExpiresAt <= 0 {
		claims.ExpiresAt = time.Now().Add(24 * time.Hour).Unix()
	}

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
func ParseRAS256TokenFileName(token string, priKeyName, pubKeyName string) (*CustomClaims, error) {
	// 私钥
	priKey, err := ioutil.ReadFile(priKeyName)
	if err != nil {
		return nil, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(priKey)
	if err != nil {
		return nil, err
	}
	// 公钥
	pubKey, err := ioutil.ReadFile(pubKeyName)
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)
	if err != nil {
		return nil, err
	}
	return ParseRAS256TokenPubKey(token, privateKey, publicKey)
}

// ParseRAS256TokenFileName 解析token
func ParseRAS256TokenPubKey(token string, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (*CustomClaims, error) {
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
		// 解密 数据
		bytes, err := DecryptWithPrivateKey(claims.Data, privateKey)
		if err != nil {
			return nil, err
		}
		claims.Data = bytes
		return claims, nil
	} else {
		return nil, errors.New("token Claims is not CustomClaims")
	}
}
