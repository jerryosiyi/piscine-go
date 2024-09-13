package piscine

func SplitWhiteSpaces(s string) []string {
	arr := []string{}
	word := ""
	for i, r := range s {
		if r == ' ' && word != "" {
			arr = append(arr, word)
			word = ""
		} else if i == len(s)-1 {
			word += string(r)
			arr = append(arr, word)
			word = ""
		} else {
			if r != ' ' {
				word += string(r)
			}
		}
	}
	return arr
}
