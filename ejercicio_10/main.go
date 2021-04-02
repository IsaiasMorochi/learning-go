package main

import (
	"fmt"
	"time"

	us "./user"
)

/*
Estructuras POO en GO
*/

// herencia, tipo polimorfismo
type person struct {
	us.Users
}

func main() {
	// User := new(user)
	// User.Id = 10
	// User.Name = "Isaias"
	// fmt.Println(User)

	u := new(person)
	u.AddUser(1, "Isaias", time.Now(), true)
	fmt.Println(u.Users)

}
