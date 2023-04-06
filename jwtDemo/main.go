package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var nowDate = time.Now().Format("2006-01-02 15")
var secret = fmt.Sprintf("%v %v", nowDate, "加盐")

// GenerateToken 生成Token值
func GenerateToken(mapClaims jwt.MapClaims, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return token.SignedString([]byte(key))
}

// token: "eyJhbGciO...解析token"
func ParseToken(token string, secret string) (jwt.MapClaims, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return claim.Claims.(jwt.MapClaims), nil
}

func main() {

	dict := make(map[string]interface{})
	dict["name"] = "张三"
	dict["age"] = 18
	tokenNew, _ := GenerateToken(dict, secret)
	fmt.Println(tokenNew, "生成token")
	q, _ := ParseToken(tokenNew, secret)
	fmt.Println(q, "解析token")
}
