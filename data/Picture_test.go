package data

import (
	"GinWebPhoto/util"
	"fmt"
	"testing"
)

func TestGetPictureListFromDir(t *testing.T) {
	var d []Picture
	d = GetPictureListFromDir("Frankcox")
	fmt.Println(d)
}

func TestSavePictureListIntoDataBase(t *testing.T) {
	util.InitDB()
	var d []Picture
	d = GetPictureListFromDir("Frankcox")
	fmt.Println(d)
	SavePictureListIntoDataBase(d)
}

func TestDeletePictureFromDir(t *testing.T) {
	username := "Frankcox"
	name := "vmw3e3.png"
	DeletePictureFromDir(username, name)
}
