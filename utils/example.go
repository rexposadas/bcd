package utils

// structs <-- objects java and c++
// objects: data structures which has fields and methods

type M struct {
}

// TODO: make the functions below as receivers for struct M
// TODO: modify the main.go file accordingly.

func Concat(strings []string) string {
	concatenated := ""
	for a, str := range strings {
		concatenated += str
		if a < len(strings)-1 {
			concatenated += " "
		}
	}
	return concatenated
}
func Add(list []int) int {
	sum := 0
	for _, number := range list {
		sum += number
	}
	return sum
}
