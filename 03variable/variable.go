package main

import "fmt"

func main() {
	// var name string = "chai aur golang"
	var name = "chai pe golang" // automatically infer type
	fmt.Println(name)

	// bool variable
	var isAdmin bool = true
	fmt.Println(isAdmin)

	// integer variable
	var age int = 18
	fmt.Println(age)

	// shorhand syntax
	programmingLng := "golang"
	fmt.Println(programmingLng)

	// variable is defined but not assign then you need to give its type
	var userId string
	userId = "12726rhryr4785"

	fmt.Println(userId)
}
