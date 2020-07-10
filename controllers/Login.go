package controllers

import (
	"GinWebPhoto/data"
	"GinWebPhoto/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//Get Login页面
func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

//Post Login页面
func LoginPost(c *gin.Context) {
	phoneNumber := c.PostForm("mobile")
	password := c.PostForm("psd")

	/*
	1.从数据库中读取，并验证是否存在该用户
	*/
	var user data.User
	result := util.Db.QueryRow("select * from User where Phone = ?",phoneNumber)
	err := result.Scan(&user.Uuid,&user.UserName,&user.UserPhone,&user.Password)
	if err != nil {
		log.Println("查询失败")
		panic(err)
	}

	real_password := user.Decode()

	if real_password  == password{
		c.SetCookie("user","123",1*3600,"/","localhost",false,true)
		c.Redirect(http.StatusMovedPermanently,"/user/homepage/"+user.UserName)
	}else {

	}

	c.SetCookie("user","123",10*60,"/","localhost",false,true)
	c.Redirect(http.StatusMovedPermanently,"/user/homepage/Frankcox")
}

//LoginWrong 页面
func LoginWrong(c *gin.Context) {
	c.HTML(http.StatusOK, "loginWrong.html", gin.H{})
	time.Sleep(3 * time.Millisecond)
	c.Redirect(http.StatusMovedPermanently, "/login")
}
