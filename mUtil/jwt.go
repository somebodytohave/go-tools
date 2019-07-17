package mUtil

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sun-wenming/go-tools/mLog"
	"github.com/sun-wenming/go-tools/mEncryptUtil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("JwtSecret")

// Claims 声明
type Claims struct {
	LoginName []byte `json:"loginname"`
	jwt.StandardClaims
}

// GenerateToken 生成 token
func GenerateToken(loginName string) (string, error) {
	var err error
	aesLoginName, err := mEncryptUtil.AesEncrypt([]byte(loginName))
	if err != nil {
		return "", err
	}

	// 现在的时间
	nowTime := time.Now()
	// 过期的时间
	expireTime := nowTime.Add(3 * time.Hour)
	// 初始化 声明
	claims := Claims{
		aesLoginName, jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "aims",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 获取完整签名之后的 token
	return tokenClaims.SignedString(jwtSecret)
}

// ParseToken 解析 token
func ParseToken(c *gin.Context) (*Claims, error) {
	token := c.Request.Header.Get("jwtToken")
	return ParseToken2(token)
}

// ParseToken2 解析 token
func ParseToken2(token string) (*Claims, error) {
	tokenClaims, _ := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("Token解析失败")
}

// GetTokenLoginName 根据 token 获取用户登录，用于去数据库获取用户id
func GetTokenLoginName(c *gin.Context) (string, error) {
	claims, err := ParseToken(c)
	if err != nil {
		mLog.GetLogger().Error(err)
		return "", errors.New("Token解析失败")
	}
	aesLoginName, err := mEncryptUtil.AesDecrypt(claims.LoginName)
	if err != nil {
		mLog.GetLogger().Error(err)
		return "", errors.New("Token解析失败")
	}
	username := string(aesLoginName)
	return username, nil
}
