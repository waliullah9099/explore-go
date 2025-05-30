package main

import "fmt"

type User struct {
	Name string // member variable or property
	Age  int
}

func printUserDetails(usr User) {
	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
}

func daw1(ccc User) {
	fmt.Println(ccc.Name)
	fmt.Println(ccc.Age)
}

func (ccc User) daw2() {
	fmt.Println(ccc.Name)
	fmt.Println(ccc.Age)
}

func (ccc User) daw3(address string) {
	fmt.Println(ccc.Name)
	fmt.Println(ccc.Age)
	fmt.Println("Tahar bari:", address)
}

// receiver function only work in custom type
func (usr User) printUser() {
	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
}

func (usr User) call(a int) {
	fmt.Println(usr.Name)
	fmt.Println(a)
}

func main() {

	/*

		user1 := User{ // instantiate
			Name: "Habib",
			Age:  32,
		}

		printUserDetails(user1)

		user1.printUser()
		user1.call(10)

	*/

	rez := User{
		Name: "Rezwan",
		Age:  4,
	}

	daw1(rez)
	rez.daw2()
	rez.daw3("Jethua")

}
