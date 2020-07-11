package middleware

import (
	"GinWebPhoto/util"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
Cookie中存放UserID,登录时与Redis中的用户匹配，如果不存在，则跳转到登录页面
*/

//VerifyCookie 验证Cookie,判断当前用户是否登录
func VerifyCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从连接池中获得连接
		var conn = util.Pool.Get()
		defer conn.Close()

		userID, err := c.Cookie("userID")
		fmt.Println(userID)
		if err != nil || userID == "" {
			c.Abort()
			c.Redirect(http.StatusMovedPermanently, "/")
		}

		verifyResult, err := redis.Bool(conn.Do("SISMEMBER", "userExists", userID))
		fmt.Println(verifyResult)
		if verifyResult != true {
			c.Abort()
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			c.Next()
		}
	}
}
