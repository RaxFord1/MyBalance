package secret

import "strings"

const mask = "*"

func maskString(input string, start int, end int, mask string) string {
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
		return maskString(input, 4, length-4, mask)
	} else if length > 4 {
		return maskString(input, 0, length-2, mask)
	} else if length > 2 {
		return maskString(input, 1, length-1, mask)
	}
	return strings.Repeat(mask, len(input))
}
