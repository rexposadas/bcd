package utils

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
