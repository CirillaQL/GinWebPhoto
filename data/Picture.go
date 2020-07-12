package data

import (
	"GinWebPhoto/util"
	_ "database/sql"
	"io/ioutil"
	"log"
)

type Picture struct {
	Name         string
	LocalAddress string
	WebAddress   string
	Describe     string
	Owner        string
}

//GetPictureListFromDir 从文件夹中读取图片信息
func GetPictureListFromDir(username string) []Picture {
	var PictureList []Picture
	srcDir := "../storage/" + username + "/Photo/"
	file, _ := ioutil.ReadDir(srcDir)
	for _, r := range file {
		Path := srcDir[1:] + r.Name()
		picture := Picture{
			Name:         r.Name(),
			LocalAddress: Path,
			WebAddress:   Path[1:],
			Describe:     "描述",
			Owner:        username,
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
		_, err := stmt.Exec(p.Name, p.LocalAddress, p.WebAddress, p.Describe, p.Owner)
		if err != nil {
			log.Fatal("插入数据库失败", err)
		}
	}
}

//GetPictureListFromDB 从数据库中读取图片信息
func GetPictureListFromDB(username string) []Picture {
	var PictureList []Picture
	var picture Picture

	rows, err := util.Db.Query("SELECT * from Picture where Owner = ?", username)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&picture.Name, &picture.LocalAddress, &picture.WebAddress, &picture.Describe, &picture.Owner)
		if err != nil {
			log.Fatal(err)
		}
		PictureList = append(PictureList, picture)
	}

	return PictureList
}

//SavePictureToDataBase 保存图片到数据库
func SavePictureToDataBase(picture Picture) (bool, error) {
	stmt, _ := util.Db.Prepare("insert into Picture(Name, LocalAddress, WebAddress, PhotoInfo, Owner) value (?,?,?,?,?);")
	defer stmt.Close()
	_, err := stmt.Exec(picture.Name, picture.LocalAddress, picture.WebAddress, picture.Describe, picture.Owner)
	if err != nil {
		log.Println("插入数据库失败！ ERROR: ", err)
		return false, err
	}
	return true, nil
}

//DeletePictureFromDB 从数据库中删除对应的图片
func DeletePictureFromDB(picture string) (bool, error) {
	stmt, err := util.Db.Prepare("DELETE FROM Picture WHERE Name = ?")
	if err != nil {
		log.Fatalln(err)
		return false, err
	}
	res, err := stmt.Exec(picture)
	_, err = res.RowsAffected()
	if err != nil {
		return false, err
	}
	return true, nil
}
