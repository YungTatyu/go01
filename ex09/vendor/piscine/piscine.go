package piscine

const (
	SW_START = iota
	SW_FIRST_DIGIT
	SW_AFTER_FIRST_DIGIT
)

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func BasicAtoi2(s string) int {
	var num int = 0
	state := SW_FIRST_DIGIT
	for _, cur_val := range s {
		switch state {
		case SW_FIRST_DIGIT:
			if !isDigit(cur_val) {
				return 0
			}
			if cur_val == '0' {
				continue
			}
			num = int(cur_val - '0')
			state = SW_AFTER_FIRST_DIGIT
		case SW_AFTER_FIRST_DIGIT:
			if !isDigit(cur_val) {
				return 0
			}
			num = 10*num + int(cur_val-'0')
		}
	}
	return num
}
