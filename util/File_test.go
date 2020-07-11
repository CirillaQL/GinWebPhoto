package util

import (
	"fmt"
	"testing"
)

func TestCreateDir(t *testing.T) {
	CreateDir("jdosadn")
}

func TestPathExists(t *testing.T) {
	result, err := PathExists("cdcds")
	fmt.Println(result)
	fmt.Println(err)
}
