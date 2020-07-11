package util

import (
	"log"
	"os"
)

func CreateDir(UserName string) {
	DirPath := "../storage/" + UserName + "/Photo"
	err := os.MkdirAll(DirPath, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}

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
