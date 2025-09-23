// comment
package main

import "fmt"

func main() {
	println("Hello World")
	name := 42
	fmt.Printf("Hello %s\n", name) // wrong: %s expects a string, got int
	fmt.Printf("%s\n", "Hello")
	for range 10 {
		fmt.Println("Hello")
	}
}
