package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtCustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	Uid   string     `json:"uid"`
	Admin bool       `json:"admin"`
	Time  *time.Time `json:"time"`
}

const key = "sample hmac secret key"

func main() {
	t := time.Now()
	a, e := createToken([]byte(key), "ssssssss",
		"123456789", true, &t)
	c, d := ParseToken(a, []byte("secret"))
	fmt.Println(e, c, d)
}

func createToken(secreKey []byte, isSuer string, uid string, isAdmin bool,
	now *time.Time) (token string, err error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
			Issuer:    isSuer,
		},
		uid,
		isAdmin,
		now,
	}
	data := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = data.SignedString([]byte("secret"))
	return
}
func ParseToken(tokenSrt string, SecretKey []byte) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims = token.Claims
	return
}
