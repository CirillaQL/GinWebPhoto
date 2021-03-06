package data

import (
	"GinWebPhoto/util"
	"encoding/base64"
	"fmt"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"log"
)

//User 用户数据库
type User struct {
	Uuid      string
	UserName  string
	UserPhone string
	Password  string
}

//Encode 将该用户的密码加密
func (user *User) Encode() {
	input := []byte(user.Password)
	encodeString := base64.StdEncoding.EncodeToString(input)
	user.Password = encodeString
}

//Decode 返回该用户的密码解密值
func (user *User) Decode() string {
	fmt.Println(user.Password)
	decodeBytes, err := base64.StdEncoding.DecodeString(user.Password)
	if err != nil {
		log.Fatalln(err)
	}
	return string(decodeBytes)
}

//CreateID 生成uuid
func CreateID() string {
	u1 := uuid.NewV4()
	uid := u1.String()
	return uid
}

//InsertUserIntoDb 向数据库插入User
func (user User) InsertUserIntoDB() (bool, error) {
	tx, _ := util.Db.Begin()
	defer tx.Rollback()
	stmt, err := util.Db.Prepare("INSERT INTO GinWebPhoto.User(uuid, UserName, Phone, Password) VALUES(?,?,?,?)")
	if err != nil {
		return false, err
	}
	result, err := stmt.Exec(user.Uuid, user.UserName, user.UserPhone, user.Password)
	if err != nil {
		return false, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return false, err
	}
	err = tx.Commit()
	return true, err
}

//InsertUserIntoRedis 向Redis中添加用户
func (user User) InsertUserIntoRedis(token string) (bool, error) {
	var conn = util.Pool.Get()
	defer conn.Close()
	fmt.Println("Token: ", token)
	repl, err := redis.Int64(conn.Do("sadd", "LoginUser", token))
	if err != nil {
		log.Println(err, repl)
		return false, err
	} else {
		_, err = conn.Do("EXPIRE", "LoginUser", 300)
		if err != nil {
			log.Fatal("用户: ", user.UserName, " 添加cookie失败！ Error: ", err)
			return false, err
		}
		return true, nil
	}
}

//CheckUserByPhone 通过Phone查询用户
func CheckUserByPhone(user *User, phone string, ch chan bool) {
	tx, _ := util.Db.Begin()
	defer tx.Rollback()
	result := util.Db.QueryRow("select * from User where Phone = ?", phone)
	err := result.Scan(&user.Uuid, &user.UserName, &user.UserPhone, &user.Password)
	if err != nil {
		tx.Rollback()
		ch <- false
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		ch <- false
	}
	ch <- true
}
