package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

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
	// claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = subject

	fmt.Println(claims["exp"])
	fmt.Println(claims["iat"])
	token.Claims = claims

	tokenString, _ := token.SignedString([]byte("testkey"))
	fmt.Println(tokenString)
	return tokenString
}

func RefreshToken(jwtToken *jwt.Token) (signedStr string) {
	//jwtToken.Claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	//currentClaims := jwtToken.Claims
	//currentClaims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()

	claim, err := jwtToken.Claims.(jwt.MapClaims)
	if !err {
		//err = errors.New("cannot convert claim to mapclaim")
		fmt.Println("cannot convert claim to mapclaim")
		return
	}
	fmt.Println("--------------------------------")
	fmt.Println(claim["exp"])
	fmt.Println(claim["iat"])
	fmt.Println(claim["sub"])
	subject := claim["sub"]
	fmt.Println("................................")
	return JwtSign(subject.(string))

}
