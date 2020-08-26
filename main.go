package main

import (
	"GinWebPhoto/controllers"
	"GinWebPhoto/middleware"
	"GinWebPhoto/util"
	"github.com/gin-gonic/gin"
	_ "github.com/gookit/color"
	"net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	util.InitDB()
	util.InitRedisPool()

	//加载静态文件
	//1.html模板
	router.LoadHTMLGlob("template/*")
	router.Static("/user/static/css", "./static/css")
	router.Static("/user/static/img", "./static/img")
	router.Static("/user/static/libs", "./static/libs")
	router.Static("/user/static/js", "./static/js")
	router.Static("/storage", "./storage")
	router.StaticFile("/favicon.ico", "./static/icon/favicon.ico")

	//未登录也可以访问的部分
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "MainPage.html", gin.H{})
	})
	router.GET("/login", controllers.LoginGet)
	router.POST("/login", controllers.LoginPost)
	router.GET("/register", controllers.RegisterGet)
	router.POST("/register", controllers.RegisterPost)
	router.GET("/quit", controllers.QuitLogin)
	//用户主界面
	user := router.Group("/user")
	user.Use(middleware.VerifyCookie())
	{
		user.GET("/homepage/:username", controllers.UserIndex)
		user.GET("/storage/:username/Photo/:img", controllers.PictureShow)
		user.GET("/homepage/:username/AddPicture", controllers.GetAddPicture)
		user.POST("/action/:username/SavePicture", controllers.GetPicture)
		user.GET("/action/:username/DeletePicture/:picture", controllers.DeletePicture)
		//user.GET("/quit", controllers.QuitLogin)
	}

	router.Run(":9090")
}
