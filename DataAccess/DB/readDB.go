package DB

import (
	"fmt"
	"go-prog/DataAccess/File"
)

func ReadSQLDB() {
	fmt.Println("ReadSQLDB called...")
	File.ReadCSVFile()
}
