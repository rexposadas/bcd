package main

import (
	"fmt"

	"github.com/rexposadas/bcd/utils"
)

func Add(x, y int) int {
	return x + y
}

func concat(list []string) {
	m := utils.M{}
	concatenated := m.Concat(list)
	fmt.Println(concatenated)
}

func main() {
	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m := utils.M{}
	a := m.Add(numberList)
	fmt.Println("sum is", a, ".")

	stringList := []string{"hello", "world", "this", "is", "a", "list", "of", "strings"}
	concat(stringList)
}
