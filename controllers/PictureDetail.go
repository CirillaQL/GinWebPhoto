package controllers

import (
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
		"img": "/storage/" + username + "/Photo/" + picture,
	})
}