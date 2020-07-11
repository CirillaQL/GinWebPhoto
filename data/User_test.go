package data

import (
	"GinWebPhoto/util"
	"fmt"
	"testing"
)

func TestUser_InsertUser(t *testing.T) {
	util.InitDB()
	j := User{
		Uuid:      CreateID(),
		UserName:  "cds",
		UserPhone: "1584061335128",
		Password:  "ccccccccdc",
	}
	j.Encode()
	fmt.Println(j)
	fmt.Println(j.CheckUser())
	for i := 0; i < 5000; i++ {
		j.InsertUserIntoDB()
	}
	//j.InsertUserIntoDB()
}

func TestUser_CheckUser(t *testing.T) {
	util.InitDB()
	j := User{
		Uuid:      "1d32dd26-b845-4108-bb9f-f156cbxsa",
		UserName:  "cdsxxsa",
		UserPhone: "1584061335xax",
		Password:  "ccccccccdc",
	}
	j.Encode()
	fmt.Println(j)
	fmt.Println(j.CheckUser())
}

func BenchmarkUser_CheckUser(b *testing.B) {
	util.InitDB()
	j := User{
		Uuid:      "1d32dd26-b845-4108-bb9f-f156cbxsa",
		UserName:  "cdsxxsa",
		UserPhone: "1584061335xax",
		Password:  "ccccccccdc",
	}
	j.Encode()
	fmt.Println(j)
	fmt.Println(j.CheckUser())
}
