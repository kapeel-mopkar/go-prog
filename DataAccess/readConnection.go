package DataAccess

import "fmt"

func init() {
	fmt.Println("Data Access/readConnection.go initialized")
}

//Exported outside package
func ReadConnection() {
	fmt.Println("Read DB connection...")
}

//Available within package
func readConfig() string {
	return "readConfig : Test Data"
}
