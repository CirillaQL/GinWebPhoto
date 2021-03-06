package controllers

import (
	"GinWebPhoto/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//从Redis中删除登录记录
func RemoveUserIDFromRedis(userID string) (bool, error) {
	var conn = util.Pool.Get()
	defer conn.Close()
	fmt.Println("当前的UserID: ", userID)
	_, err := conn.Do("srem", "userExists", userID)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//退出登录
func QuitLogin(c *gin.Context) {
	fmt.Println("调用退出")

	userID, err := c.Cookie("userID")
	if nil != err {
		c.Redirect(http.StatusMovedPermanently, "http://localhost:9090")
	}
	c.SetCookie("userID", "", -1, "/", "localhost", false, true)
	result, err := RemoveUserIDFromRedis(userID)
	if result == false || err != nil {
		log.Fatalln("在Redis中删除用户ID失败! 用户ID：", userID, "  Error: ", err)
	} else {
		c.Redirect(http.StatusMovedPermanently, "http://localhost:9090")
	}

}
