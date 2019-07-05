package utils

func IsBlank(str string) bool {
	if str == "" || len(str) == 0 {
		return true
	}
	for _, c := range str {
		switch c {
		case ' ', '\t', '\r', '\n':
		default:
			return false
		}
	}
	return true
}
