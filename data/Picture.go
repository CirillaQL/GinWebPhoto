package data

import (
	"GinWebPhoto/util"
	"log"
	"os"
	"time"
)

//Picture 图片的结构体
type Picture struct {
	Id           string
	Name         string
	LocalAddress string
	WebAddress   string
	Describe     string
	CreateTime   time.Time
	Owner        string
}

//GetPictureListFromDir 从文件夹中读取图片信息
/*
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
*/
//SavePictureListIntoDataBase 将本地文件存入数据库
func SavePictureListIntoDataBase(picture []Picture) {
	stmt, err := util.Db.Prepare("INSERT INTO Picture(id,NAME, LocalAddress, WebAddress, PhotoInfo,CreateTime, Owner) VALUES (?,?,?,?,?)")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	for _, p := range picture {
		_, err := stmt.Exec(p.Id, p.Name, p.LocalAddress, p.WebAddress, p.Describe, p.CreateTime, p.Owner)
		if err != nil {
			log.Fatal("插入数据库失败", err)
		}
	}
}

//GetPictureListFromDB 从数据库中读取图片信息
func GetPictureListFromDB(username string) []Picture {
	var PictureList []Picture
	var picture Picture

	tx, err := util.Db.Begin()
	if err != nil {
		log.Println("从数据库中读取图片列表时出错： Error", err)
		tx.Rollback()
	}
	defer tx.Commit()
	rows, err := tx.Query("SELECT * from Picture where Owner = ?", username)
	if err != nil {
		log.Println("从数据库中读取图片列表时出错： Error", err)
		tx.Rollback()
	}
	for rows.Next() {
		err = rows.Scan(&picture.Id, &picture.Name, &picture.LocalAddress, &picture.WebAddress, &picture.Describe, &picture.CreateTime, &picture.Owner)
		if err != nil {
			log.Fatal(err)
			tx.Rollback()
		}
		PictureList = append(PictureList, picture)
	}

	return PictureList
}

//SavePictureToDataBase 保存图片到数据库
func SavePictureToDataBase(picture Picture) (bool, error) {
	//使用事务进行操作，
	tx, _ := util.Db.Begin()
	defer tx.Rollback()
	_, err := tx.Exec("insert into Picture(id,Name, LocalAddress, WebAddress, PhotoInfo,CreateTime ,Owner) value (?,?,?,?,?,?,?);",
		picture.Id, picture.Name, picture.LocalAddress, picture.WebAddress, picture.Describe, picture.CreateTime, picture.Owner)
	if err != nil {
		log.Println("插入数据库失败！ ERROR: ", err)
		_ = tx.Rollback()
		return false, err
	}
	err = tx.Commit()
	if err != nil {
		log.Println("插入数据库失败！ ERROR: ", err)
		_ = tx.Rollback()
		return false, err
	}
	return true, nil
}

//DeletePictureFromDB 从数据库中删除对应的图片
func DeletePictureFromDB(picture string) (bool, error) {
	tx, _ := util.Db.Begin()
	defer tx.Rollback()
	stmt, err := util.Db.Prepare("DELETE FROM Picture WHERE Name = ?")
	if err != nil {
		log.Fatalln(err)
		return false, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(picture)
	_, err = res.RowsAffected()
	err = tx.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}

//DeletePictureFromDir 从本地硬盘中删除对应的
func DeletePictureFromDir(username string, filename string) (bool, error) {
	path := "./storage/" + username + "/Photo/" + filename
	err := os.Remove(path)
	if err != nil {
		return false, err
	}
	return true, nil
}
