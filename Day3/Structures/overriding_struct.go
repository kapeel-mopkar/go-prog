package main

import "fmt"

type User1 struct {
	Name string
	ID   int
}

type Admin1 struct {
	User1 // Embed, variables and methods of embedded struct
	Role  string
}

func (user User1) GetDetails() {
	fmt.Println("\tGetDetails() for User - ", user.ID, user.Name)
}

func (admin Admin1) GetDetails() {
	fmt.Println("\tGetDetails() for Admin - ", admin.ID, admin.Name, admin.Role)
}

func main() {
	var user1 = User1{"Kapeel", 123}
	var admin1 = Admin1{user1, "AdminRole"}

	fmt.Println(admin1.ID, admin1.Name, admin1.Role)
	user1.GetDetails()
	admin1.GetDetails()

	admin2 := Admin1{
		User1: User1{
			Name: "Sam",
			ID:   224,
		}, Role: "Moderator",
	}

	fmt.Println(admin2.ID, admin2.Name, admin2.Role)
	admin2.GetDetails()

}
