package controllers

import (
	"GinWebPhoto/data"
	"GinWebPhoto/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//Get命令时Handler
func RegisterGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

/*
1.从Post解析数据
2.生成用户结构体
3.密码加密
4.插入到数据库
5.生成文件夹
*/
//RegisterPost 用户注册处理
func RegisterPost(c *gin.Context) {
	username := c.PostForm("name")
	mobile := c.PostForm("mobile")
	password := c.PostForm("psd")

	log.Println("用户注册:    ", "用户名：", username, "手机号码：", mobile, "密码：", password)
	userRegister := data.User{
		Uuid:      data.CreateID(),
		UserName:  username,
		Password:  password,
		UserPhone: mobile,
	}
	userRegister.Encode()

	result, err := userRegister.InsertUserIntoDB()
	if err != nil || result != true {
		log.Println("注册错误")
		log.Println(err)
		log.Println(result)
	}
	util.CreateDir(username)
	if err != nil || result != true {
		log.Println("注册错误")
		log.Println(err)
		log.Fatalln(result)
	}
	//生成Token
	token, err := util.GenerateToken(userRegister.Uuid, userRegister.UserName, userRegister.Password)
	if err != nil {
		log.Fatalln("Token生成错误")
	}
	result, err = userRegister.InsertUserIntoRedis(token)
	if result == false && err != nil {
		log.Fatalln("Redis存入失败: error: ", err)
	}
	c.SetCookie("token", token, 300, "/", "localhost", false, true)
	c.Redirect(http.StatusMovedPermanently, "/user/homepage/"+username)
}
