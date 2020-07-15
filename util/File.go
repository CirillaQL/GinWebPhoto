package util

import (
	"log"
	"os"
)

//CreateDir 创建文件夹
func CreateDir(UserName string) {
	DirPath := "../storage/" + UserName + "/Photo"
	err := os.MkdirAll(DirPath, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}

//PathExists 判断文件夹是否存在
func PathExists(username string) (bool, error) {
	path := "../storage/" + username
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
