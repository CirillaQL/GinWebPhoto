package controllers

import (
	"GinWebPhoto/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

	data.DeletePictureFromDB(picture)

}
