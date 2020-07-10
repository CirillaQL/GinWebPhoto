package data

import (
	"GinWebPhoto/util"
	"encoding/base64"
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
func (user User) Encode(){
	input := []byte(user.Password)
	encodeString := base64.StdEncoding.EncodeToString(input)
	user.Password = encodeString
}

//Decode 返回该用户的密码解密值
func (user User) Decode() string{
	decodeBytes, err := base64.StdEncoding.DecodeString(user.Password)
	if err != nil {
		log.Fatalln(err)
	}
	ans, err := base64.StdEncoding.DecodeString(string(decodeBytes))
	if err != nil {
		log.Fatalln(err)
	}
	return string(ans)
}

//CreateID 生成uuid
func CreateID() string {
	u1 := uuid.NewV4()
	uid := u1.String()
	return uid
}

//InsertUserIntoDb 向数据库插入User
func (user User) InsertUserIntoDB() (bool,error){
	stmt, err := util.Db.Prepare("INSERT INTO GinWebPhoto.User(uuid, UserName, Phone, Password) VALUES(?,?,?,?)")
	if err != nil{
		return false,err
	}
	result, err := stmt.Exec(user.Uuid,user.UserName,user.UserPhone,user.Password)
	if err != nil{
		return false,err
	}
	result_id,err := result.LastInsertId()
	if err != nil || result_id == 0{
		return false,err
	}
	return true,err
}

//InsertUserIntoRedis 向Redis中添加用户
func (user User) InsertUserIntoRedis() (bool,error){
	var conn = util.Pool.Get()
	defer conn.Close()

	repl ,err := redis.Int64(conn.Do("sadd", "userExists", user.Uuid))
	if err != nil || repl != 1{
		return false,err
	}else {
		return true,err
	}
}

