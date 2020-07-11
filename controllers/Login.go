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
	result := util.Db.QueryRow("select * from User where Phone = ?", phoneNumber)
	err := result.Scan(&user.Uuid, &user.UserName, &user.UserPhone, &user.Password)
	fmt.Println(user)
	if err != nil {
		log.Println("查询失败")
		c.Redirect(http.StatusMovedPermanently, "/")
	}

	realPassword := user.Decode()
	fmt.Println("输入密码: ", password)
	fmt.Println("用户密码: ", realPassword)
	if realPassword == password {
		result, err := user.InsertUserIntoRedis()
		fmt.Println("redis")
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
			c.SetCookie("userID", user.Uuid, 1*3600, "/", "localhost", false, true)
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
