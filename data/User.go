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
	return true, err
}

//InsertUserIntoRedis 向Redis中添加用户
func (user User) InsertUserIntoRedis() (bool, error) {
	var conn = util.Pool.Get()
	defer conn.Close()

	repl, err := redis.Int64(conn.Do("sadd", "userExists", user.Uuid))
	if err != nil {
		log.Println(err, repl)
		return false, err
	} else {
		_, err = conn.Do("EXPIRE", "userExists", 3600)
		if err != nil {
			log.Fatal("用户: ", user.UserName, " 添加cookie失败！ Error: ", err)
			return false, err
		}
		return true, nil
	}
}

//CheckUserByID 通过ID查询用户
func (user User) CheckUserByID(ch chan bool) {
	var result string
	rows, err := util.Db.Query("SELECT uuid FROM User;")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			log.Fatal(err.Error())
		}
		if user.Uuid == result {
			ch <- false
		} else {
			continue
		}
	}
	ch <- true
}

//CheckUserByName 通过Name查询用户
func (user User) CheckUserByName(ch chan bool) {
	var result string
	rows, err := util.Db.Query("SELECT UserName FROM User;")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			log.Fatal(err.Error())
		}
		if user.UserName == result {
			ch <- false
		} else {
			continue
		}
	}
	ch <- true
}

//CheckUserByPhone 通过Phone查询用户
func (user User) CheckUserByPhone(ch chan bool) {
	var result string
	rows, err := util.Db.Query("SELECT Phone FROM User;")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			log.Fatal(err.Error())
		}
		if user.Uuid == result {
			ch <- false
		} else {
			continue
		}
	}
	ch <- true
}

//CheckUser 并行调用三个检查，判断数据库中是否有重名
func (user User) CheckUser() bool {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	go user.CheckUserByID(ch1)
	go user.CheckUserByName(ch2)
	go user.CheckUserByPhone(ch3)
	result_1 := <-ch1
	result_2 := <-ch2
	result_3 := <-ch3
	if result_1 == true && result_2 == true && result_3 == true {
		return true
	} else {
		return false
	}
}
