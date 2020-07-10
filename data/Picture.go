package data

import (
	"GinWebPhoto/util"
	"io/ioutil"
	"log"
)

type Picture struct {
	Name string
	LocalAddress string
	WebAddress string
	Describe string
	Owner string
}

//GetPictureListFromDir 从文件夹中读取图片信息
func GetPictureListFromDir(username string) []Picture {
	var PictureList []Picture
	srcDir := "../storage/" + username + "/Photo/"
	file, _ := ioutil.ReadDir(srcDir)
	for _, r := range file {
		Path := srcDir[1:] + r.Name()
		picture := Picture{
			Name:     r.Name(),
			LocalAddress:  Path,
			WebAddress: Path,
			Describe: "描述",
			Owner:    username,
		}
		PictureList = append(PictureList, picture)
	}
	return PictureList
}

//SavePictureListIntoDataBase 将本地文件存入数据库
func SavePictureListIntoDataBase(picture []Picture) {
	stmt, err := util.Db.Prepare("INSERT INTO Picture(NAME, LocalAddress, WebAddress, PhotoInfo, Owner) VALUES (?,?,?,?,?)")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	for _, p := range picture {
		_, err := stmt.Exec(p.Name, p.LocalAddress, p.WebAddress,p.Describe, p.Owner)
		if err != nil {
			log.Fatal("插入数据库失败", err)
		}
	}
}

