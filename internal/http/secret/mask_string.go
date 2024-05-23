package secret

import "strings"

func MaskString(input string, start int, end int, mask string) string {
	if end >= len(input) || start > len(input) {
		return "****"
	}
	if start > end {
		tmp := start
		start = end
		end = tmp
	}

	return input[:start] + strings.Repeat(mask, end-start) + input[end:]
}

func ApplyMask(input string) string {
	length := len(input)
	if length > 8 {
		return MaskString(input, 4, length-4, "*")
	} else if length > 4 {
		return MaskString(input, 0, length-2, "*")
	}
	return "*"
}
