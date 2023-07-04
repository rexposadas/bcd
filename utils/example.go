package utils

type M struct {
}

func (m M) Concat(strings []string) string {
	concatenated := ""
	for a, str := range strings {
		concatenated += str
		if a < len(strings)-1 {
			concatenated += " "
		}
	}
	return concatenated
}

func (m M) Add(list []int) int {
	sum := 0
	for _, number := range list {
		sum += number
	}
	return sum
}

func (m M) Sub(x, y int) int {
	return x - y
}

// todo: implement
//
// test 1: 1,2,4,1,2,3
// output: 1,2
func (m M) Dupes(list []int) []int {
	return []int{}
}
