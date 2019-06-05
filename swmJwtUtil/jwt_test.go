/*
 * Genarate rsa keys.
 */

package main

import (
	"fmt"
	"github.com/sun-wenming/go-tools/swmJwtUtil"
	"log"
)

// openssl 私钥 公钥
//openssl genrsa -out rsa_private_key.pem 2048
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
func main() {

	customClaims := swmJwtUtil.CustomClaims{Data: []byte("abcde")}
	// TODO 修改私钥公钥 名称
	token, err := swmJwtUtil.GenRSA256TokenWithFileName(customClaims, "rsa_private_key.pem", "rsa_public_key.pem")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token)
	claims, err := swmJwtUtil.ParseRAS256TokenFileName(token, "rsa_private_key.pem", "rsa_public_key.pem")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("------------------")
	fmt.Println(claims.Data)
	fmt.Println(string(claims.Data))
	fmt.Println("===========")

}
