package pusher

import "strconv"

func CheckDup(a Stack) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] == a[i+1] {
			return true
		}
	}
	return false
}

// Parsing utility function
func ParseInt(s string) (int, error) {
	return strconv.Atoi(s)
}
