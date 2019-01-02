// reusable package
package DataAccess

import (
	"fmt"
)

var FileName string = "file.txt"
var FileLocation string = "c:\\projects"

func init() {
	fmt.Println("Data Access/readData.go initialized")
}

func ReadData() {
	fmt.Println("ReadData is called")
}

func WriteData() {
	var config = readConfig()
	fmt.Println("WriteData is called", config)
}
