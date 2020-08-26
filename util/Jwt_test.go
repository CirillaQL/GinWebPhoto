package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

func TestCreateToken(t *testing.T) {
	token, _ := GenerateToken("95ea83f2-5bc3-4518-afc2-9b785c4cb057", "Frankcox", "ql1194946223")
	fmt.Println(token)

	claims, err := ParseToken(token, []byte(Secret))
	if nil != err {
		fmt.Println(" err :", err)
	}
	fmt.Println("claims:", claims)
	fmt.Println("claims uid:", claims.(jwt.MapClaims)["UserID"])
}
