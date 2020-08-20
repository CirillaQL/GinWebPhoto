package controllers

import (
	"GinWebPhoto/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserIndex(c *gin.Context) {
	username := c.Param("username")
	LoadPicture := data.GetPictureListFromDB(username)
	c.HTML(http.StatusOK, "photo.html", gin.H{
		"name":  username,
		"image": LoadPicture,
	})
}
