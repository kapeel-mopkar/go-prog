package main

import (
	"go-prog/DataAccess"
	"go-prog/DataAccess/DB"
	dbCfg "go-prog/DataAccess/DB/Config" // package aliases
	"go-prog/DataAccess/File"

	//fileCfg "DataAccess/File/Config" // package aliases
	"fmt"

	"github.com/golang/example/stringutil"
)

func main() {
	fmt.Println("Hello world")
	DataAccess.ReadData()
	DataAccess.WriteData()
	DataAccess.ReadConnection()

	fmt.Println(DataAccess.FileName)
	fmt.Println(DataAccess.FileLocation)

	DB.ReadSQLDB()
	File.ReadCSVFile()

	dbCfg.ReadWebConfig()

	//fileCfg.ReadServerConfig()
	//using imported packages using go get
	fmt.Println("Reversing using StringUtil ->", stringutil.Reverse("PSL"))

	//Access function from main package
	fmt.Println("Support Data : ", GetSupportData())
}
