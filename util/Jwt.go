package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"time"
)

const (
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin    = "请重新登陆"
)

type Claims struct {
	UserID      string   `json:userID`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	Permissions []string `json:"permissions"`
	jwt.StandardClaims
}

var (
	Secret = "dong_tech"
)

//生成目标Token
func getToken(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New(ErrorReason_ServerBusy)
	}
	return signedToken, nil
}

//获取用户ID,用户名Username,用户密码Password生成对应的Token(String)序列
func GenerateToken(userid, username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(5 * time.Minute)

	claims := Claims{
		UserID:      userid,
		Username:    username,
		Password:    password,
		Permissions: []string{},
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = expireTime.Unix()

	signedToken, err := getToken(&claims)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
