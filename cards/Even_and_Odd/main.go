package main

import (
	"fmt"
)

func main() {
	numbers := []int{}
	i := 0
	for i <= 10 {
		numbers = append(numbers, i)
		i++
	}
	for _, n := range numbers {
		if n%2 == 0 {
			fmt.Println(n, " is even")
		} else {
			fmt.Println(n, " is odd")
		}
	}
}
