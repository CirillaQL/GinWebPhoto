package controllers

import (
	"GinWebPhoto/data"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//Get命令时Handler
func RegisterGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

//Post命令时Handler
func RegisterPost(c *gin.Context) {
	username := c.PostForm("name")
	mobile := c.PostForm("mobile")
	password := c.PostForm("psd")
	log.Println("用户注册:    ", "用户名：", username, "手机号码：", mobile, "密码：", password)
	/*
		生成结构体
	*/
	//生成用户
	userRegister := data.User{
		Uuid:      data.CreateID(),
		UserName:  username,
		Password:  password,
		UserPhone: mobile,
	}
	userRegister.Encode()

	result,err := userRegister.InsertUserIntoDB()
	if err != nil || result != true{
		log.Println("注册错误")
	}

}

