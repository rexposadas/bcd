package main

import (
	"fmt"
)

// todo: be able to add the numbers in the list
func add(list []int) int {
	sum := 0
	for _, number := range list {
		sum += number
	}
	return sum
}

// todo: concatenate the strings in the list
func concat(strings []string) string {
	concatenated := ""
	for a, str := range strings {
		concatenated += str
		if a < len(strings)-1 {
			concatenated += " "
		}
	}
	return concatenated
}

func main() {

	// todo: implement all this
	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a := add(numberList)
	fmt.Println("sum is", a, ".")

	stringList := []string{"hello", "world", "this", "is", "a", "list", "of", "strings"}
	concatenate := concat(stringList)
	// output: "hello world this is a list of strings"
	fmt.Println(concatenate)

}
