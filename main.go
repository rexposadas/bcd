package main

import (
	"fmt"
	"github.com/rexposadas/bcd/utils"
)


func main() {
	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a := utils.Add(numberList)
	fmt.Println("sum is", a, ".")

	stringList := []string{"hello", "world", "this", "is", "a", "list", "of", "strings"}
	concatenate := utils.Concat(stringList)
	fmt.Println(concatenate)
}
