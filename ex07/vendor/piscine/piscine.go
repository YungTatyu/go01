package piscine

func StrRev(s string) string {
	var rs string
	for _, v := range s {
		rs = string(v) + rs
	}
	return rs
}
