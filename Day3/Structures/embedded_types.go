package main

import "fmt"

type User struct {
	Name string
	ID   int
}

type Admin struct {
	User // Embed, variables and methods of embedded struct
	Role string
}

func (user User) getUserDetails() {
	fmt.Println("GetUserDetails() - ", user.ID, user.Name)
}

func main() {
	var user1 = User{"Kapeel", 123}
	var admin1 = Admin{user1, "AdminRole"}

	fmt.Println(admin1.ID, admin1.Name, admin1.Role)
	user1.getUserDetails()

	admin2 := Admin{
		User: User{
			Name: "Sam",
			ID:   224,
		}, Role: "Moderator",
	}

	fmt.Println(admin2.ID, admin2.Name, admin2.Role)

}
