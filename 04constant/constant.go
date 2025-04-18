package main

import "fmt"

// outside the main funtion
// const age = 30

// cant do that
// name:= "jhghddg"

func main() {
	const pi string = "go"
	// cant do that
	// pi = "javascript"
	fmt.Println(pi)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
