package main

import "fmt"

// for -> only construct in go for looping
func main() {
	// implement while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// infite loop
	// for {
	// 	println("1")
	// }

	// for loop
	for i := 0; i < 3; i++ {
		println(i)
	}

}
