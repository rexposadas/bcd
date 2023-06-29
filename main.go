package main

import (
	"fmt"
	"github.com/rexposadas/bcd/utils"
)

// todo: be able to add the numbers in the list
func add(x, y int) int {
	return x + y
}

// todo: concatenate the strings in the list
func concat(list []string) {

}

func main() {
	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a := utils.Add(numberList)
	fmt.Println("sum is", a, ".")

	stringList := []string{"hello", "world", "this", "is", "a", "list", "of", "strings"}
	concatenate := utils.Concat(stringList)
	fmt.Println(concatenate)
}
