package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer)
	fmt.Println(namePointer)
	printPointer(namePointer)
	printPointer(&name)

}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
	fmt.Println(namePointer)
	fmt.Println(*namePointer)
}
