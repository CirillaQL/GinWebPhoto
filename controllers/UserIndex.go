package controllers

import (
	"GinWebPhoto/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserIndex(c *gin.Context){
	username := c.Param("username")

	LoadPicture := data.GetPictureListFromDir(username)
	fmt.Println(LoadPicture)
	_test1 := data.Picture{
		Name:         "山",
		LocalAddress: "/storage/Frankcox/Photo/山.jpg",
		WebAddress:   "storage/Frankcox/Photo/山.jpg",
		Describe:     "dcscds",
		Owner:        "Frankcox",
	}
	_test2 := data.Picture{
		Name:         "漫画.jpg",
		LocalAddress: "/storage/Frankcox/Photo/漫画.jpg",
		WebAddress:   "storage/Frankcox/Photo/漫画.jpg",
		Describe:     "dcscds",
		Owner:        "Frankcox",
	}
	var k []data.Picture
	k = append(k,_test1,_test2)
	c.HTML(http.StatusOK, "photo.html", gin.H{
		"name":  username,
		"image": k,
	})
}
