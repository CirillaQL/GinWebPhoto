package controllers

import (
	"GinWebPhoto/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//GetAddPicture 获取添加图片页面
func GetAddPicture(c *gin.Context) {
	name := c.Param("username")
	c.HTML(200, "upload.html", gin.H{
		"name": name,
	})
}

//GetPicture 保存图片到本地
func GetPicture(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	des := ctx.PostForm("text")
	if err != nil {
		log.Fatal(err)
		return
	}
	username := ctx.Param("username")
	path := "./storage/" + username + "/Photo/" + file.Filename
	ok := ctx.SaveUploadedFile(file, path)
	if ok != nil {
		log.Fatal("保存时错误")
	} else {
		log.Println(file.Filename, "保存成功")
		//接下来保存到数据库
		picture := data.Picture{
			Name:         file.Filename,
			LocalAddress: path,
			WebAddress:   path[1:],
			Describe:     des,
			Owner:        username,
		}
		fmt.Println(picture)
		result, err := data.SavePictureToDataBase(picture)
		if result != true {
			panic(err)
		}
		//保存后，跳转到主页面
		ctx.Redirect(http.StatusMovedPermanently, "/user/homepage/"+username)
	}

}
