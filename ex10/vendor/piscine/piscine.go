package piscine

const (
	SW_START = iota
	SW_SIGN  = iota
	SW_FIRST_DIGIT
	SW_AFTER_FIRST_DIGIT
	SW_END
)

func StrLen(s string) int {
	var i int = 0
	for range s {
		i++
	}
	return i
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func Atoi(s string) int {
	var num int = 0
	sLen := StrLen(s)
	state := SW_START
	sign := 1 // 1 or -1
	i := 0
	for {
		if i >= sLen {
			break
		}
		var cur_val rune = rune(s[i])
		switch state {
		case SW_START:
			state = SW_SIGN
		case SW_SIGN:
			if cur_val == '+' || cur_val == '-' {
				i++
				if cur_val == '-' {
					sign = -1
				}
			}
			state = SW_FIRST_DIGIT
		case SW_FIRST_DIGIT:
			i++
			if !isDigit(cur_val) {
				return 0
			}
			if cur_val == '0' {
				continue
			}
			num = int(cur_val - '0')
			state = SW_AFTER_FIRST_DIGIT
		case SW_AFTER_FIRST_DIGIT:
			i++
			if !isDigit(cur_val) {
				return 0
			}
			num = 10*num + int(cur_val-'0')
		}
	}
	return num * sign
}
