/*
 * Genarate rsa keys.
 */

package mjwt

import (
	"fmt"
	"log"
)

const (
	priKeyName = "jwt_rsa_private_key.pem"
	pubKeyName = "jwt_rsa_public_key.pem"
)

// openssl 私钥 公钥
//openssl genrsa -out rsa_private_key.pem 2048
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
func main() {

	customClaims := CustomClaims{Data: []byte("")}
	// TODO 修改私钥公钥 名称

	token, err := GenToken("abcde")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token)
	claims, err := ParseToken(token)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("------------------")
	fmt.Println(claims.Data)
	fmt.Println(string(claims.Data))
	fmt.Println("===========")

}

func GenToken(data string) (string, error) {
	customClaims := CustomClaims{Data: []byte(data)}
	token, err := GenRSA256TokenWithFileName(customClaims, priKeyName, pubKeyName)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (*CustomClaims, error) {
	claims, err := ParseRAS256TokenFileName(token, priKeyName, pubKeyName)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
