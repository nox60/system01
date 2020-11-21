package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//func main() {
//	fmt.Println("test")
//	signedStr := JwtSign("ddddd")
//
//	//token, _ := jwt.Parse(signedStr, func(token *jwt.Token) (interface{}, error) {
//	//	// Don't forget to validate the alg is what you expect:
//	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//	//		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//	//	}
//	//
//	//	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
//	//	return []byte("testkey"), nil
//	//})
//
//	token, _ := JwtParse(signedStr)
//
//	fmt.Println("-----")
//	fmt.Println(token)
//
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		fmt.Println(claims["sub"])
//	} else {
//		fmt.Println("error")
//	}
//}

func JwtParse(signedStr string) (jwtToken *jwt.Token, err error) {
	token, err := jwt.Parse(signedStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("testkey"), nil
	})

	return token, err
}

func JwtSign(subject string) (signedStr string) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = subject
	token.Claims = claims

	tokenString, _ := token.SignedString([]byte("testkey"))

	fmt.Println(tokenString)

	return tokenString
}
