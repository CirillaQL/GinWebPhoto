package data

import (
	"GinWebPhoto/util"
	"fmt"
	"testing"
)

func TestUser_InsertUser(t *testing.T) {
	util.InitDB()
	j := User{
		Uuid:      	CreateID(),
		UserName:  "Frankcox",
		UserPhone: "15840613358",
		Password:  	"ql1194946223",
	}
	j.Encode()
	fmt.Println(j)
	j.InsertUserIntoDB()
}