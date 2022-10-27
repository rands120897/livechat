package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GetTokenJWT(email string) string {
	mySignKey := []byte("Kadal")

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		Issuer:    email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(mySignKey)
	fmt.Println(err)

	return tokenString

}

func CheckValidTokenJwt(tokenString string) bool {

	claims := jwt.StandardClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {

		return []byte("Kadal"), nil
	})

	fmt.Println(claims.Issuer)

	if token.Valid {

		fmt.Println("okee token valid")
		return true

	} else {
		fmt.Println("Token TIdak Valid :", err)
		return false
	}

}
