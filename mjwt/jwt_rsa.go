package mjwt

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Data      []byte `json:"data"`
	ExtraData []byte `json:"extra_data"`
	jwt.StandardClaims
}

// GenRSA256TokenByFile
func GenRSA256TokenByFile(claims CustomClaims, priKeyFile, pubKeyFile string) (string, error) {
	// 私钥
	privateKey, err := GetPriKey(priKeyFile)
	if err != nil {
		return "", err
	}
	// 公钥
	publicKey, err := GetPubKey(pubKeyFile)
	if err != nil {
		return "", err
	}
	return GenRSA256Token(claims, privateKey, publicKey)
}

// GenRSA256TokenByFilePwd
func GenRSA256TokenByFilePwd(claims CustomClaims, priKeyFile, pubKeyFile, keyPwd string) (string, error) {
	// 私钥
	privateKey, err := GetPriKeyPwd(priKeyFile, keyPwd)
	if err != nil {
		return "", err
	}
	// 公钥
	publicKey, err := GetPubKey(pubKeyFile)
	if err != nil {
		return "", err
	}
	return GenRSA256Token(claims, privateKey, publicKey)
}

// GenRSA256Token 生成 密钥对加密的 token
func GenRSA256Token(claims CustomClaims, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (string, error) {
	bytes, err := EncryptWithPublicKey(claims.Data, publicKey)
	if err != nil {
		return "", err
	}
	extraDataBytes, err := EncryptWithPublicKey(claims.ExtraData, publicKey)
	if err != nil {
		return "", err
	}

	claims.Data = bytes
	claims.ExtraData = extraDataBytes
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

// ParseRAS256TokenByFile
func ParseRAS256TokenByFile(token string, priKeyFile, pubKeyFile string) (*CustomClaims, error) {
	privateKey, err := GetPriKey(priKeyFile)
	if err != nil {
		return nil, err
	}
	publicKey, err := GetPubKey(pubKeyFile)
	if err != nil {
		return nil, err
	}
	return ParseRAS256TokenByKey(token, privateKey, publicKey)
}

// ParseRAS256TokenByFilePwd
func ParseRAS256TokenByFilePwd(token string, priKeyFile, pubKeyFile, keyPwd string) (*CustomClaims, error) {
	privateKey, err := GetPriKeyPwd(priKeyFile, keyPwd)
	if err != nil {
		return nil, err
	}
	publicKey, err := GetPubKey(pubKeyFile)
	if err != nil {
		return nil, err
	}
	return ParseRAS256TokenByKey(token, privateKey, publicKey)
}

// ParseRAS256TokenByKey 解析token
func ParseRAS256TokenByKey(token string, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (*CustomClaims, error) {
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
		// 解密 数据
		extraDataBytes, err := DecryptWithPrivateKey(claims.ExtraData, privateKey)
		if err != nil {
			return nil, err
		}
		claims.Data = bytes
		claims.ExtraData = extraDataBytes
		return claims, nil
	} else {
		return nil, errors.New("token Claims is not CustomClaims")
	}
}

// 获取私钥
func GetPriKey(priKeyFile string) (*rsa.PrivateKey, error) {
	priKey, err := ioutil.ReadFile(priKeyFile)
	if err != nil {
		return nil, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(priKey)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// 获取 加密码的私钥
func GetPriKeyPwd(priKeyFile, pwd string) (*rsa.PrivateKey, error) {
	priKey, err := ioutil.ReadFile(priKeyFile)
	if err != nil {
		return nil, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEMWithPassword(priKey, pwd)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// 获取公钥
func GetPubKey(pubKeyFile string) (*rsa.PublicKey, error) {
	pubKey, err := ioutil.ReadFile(pubKeyFile)
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}
