package DB

import (
	"DataAccess/File"
	"fmt"
)

func ReadSQLDB() {
	fmt.Println("ReadSQLDB called...")
	File.ReadCSVFile()
}
