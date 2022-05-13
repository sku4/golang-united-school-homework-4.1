package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		output += string(runes[len(runes)-1-i])
	}
	return output
}
