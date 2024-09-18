package poker

func reverseString(str string) string {
	bytes := []byte{}
	for i := len(str) - 1; i >= 0; i-- {
		bytes = append(bytes, str[i])
	}
	return string(bytes)
}
