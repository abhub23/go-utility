package rim

import (
	"strings"
)

func Flames(name1 string, name2 string) string {
	lower1 := strings.ToLower(name1)
	lower2 := strings.ToLower(name2)
	return lower1 + " Flames " + lower2
}
