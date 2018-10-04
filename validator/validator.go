package validator

// ValidateParens checks if the parens are ballanced in the input string.
func ValidateParens(in string) bool {
	if len(in) == 0 {
		panic("Bad input!")
	}
	x := 0
	sum := 0

	for x < len(in) {
		ch := in[x]
		switch ch {
		case '(':
			sum++
		case ')':
			sum--
			if sum < 0 {
				return false
			}
		default:
			return false
		}

		x++
	}
	return (sum == 0)
}
