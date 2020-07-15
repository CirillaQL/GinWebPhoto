package controllers

import (
	"GinWebPhoto/data"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PictureShow(context *gin.Context) {
	username := context.Param("username")
	picture := context.Param("img")
	fmt.Println(username)
	fmt.Println(picture)
	context.HTML(http.StatusOK, "PictureCheck.html", gin.H{
		"img":     "/storage/" + username + "/Photo/" + picture,
		"name":    username,
		"imgname": picture,
	})
}

func DeletePicture(context *gin.Context) {
	picture := context.Param("picture")
	username := context.Param("username")
	data.DeletePictureFromDB(picture)
	ans, err := data.DeletePictureFromDir(username, picture)
	fmt.Println(ans)
	fmt.Println(err)
	context.Redirect(http.StatusMovedPermanently, "/user/homepage/"+username)
}
