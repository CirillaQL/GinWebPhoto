package controllers

import (
	"GinWebPhoto/data"
	"GinWebPhoto/util"
	"fmt"
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
		1.解析Post请求，获取电话和密码
		2.从数据库中查找用户
		3.匹配密码
		4.匹配成功，存入Redis
		5.检查本地是否存在用户文件夹
		6.设置Cookie,跳转到主页(注意，此处应该调整Cookie保存时间，暂时按照1小时处理)
	*/
	var user data.User

	//从数据库中查询用户信息
	ch := make(chan bool, 1)
	go data.CheckUserByPhone(&user, phoneNumber, ch)

	ans := <-ch
	if ans == false {
		log.Println("登录时查询账号信息失败！")
		c.Redirect(http.StatusMovedPermanently, "/")
	}
	realPassword := user.Decode()
	if realPassword == password {
		token, err := util.GenerateToken(user.Uuid, user.UserName, user.Password)
		result, err := user.InsertUserIntoRedis(token)
		if err != nil || result == false {
			log.Fatalln(err)
		} else {
			//检查本地是否存在目标文件夹
			result, err = util.PathExists(user.UserName)
			if err != nil {
				log.Fatalln(err)
			}
			if result == false {
				util.CreateDir(user.UserName)
			}
			c.SetCookie("token", token, 1*3600, "/", "localhost", false, true)
			c.Redirect(http.StatusMovedPermanently, "/user/homepage/"+user.UserName)
		}
	} else {
		fmt.Println("Not redis")
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}

//LoginWrong 页面
func LoginWrong(c *gin.Context) {
	c.HTML(http.StatusOK, "loginWrong.html", gin.H{})
	time.Sleep(3 * time.Millisecond)
	c.Redirect(http.StatusMovedPermanently, "/login")
}
