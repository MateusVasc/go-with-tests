package iteration

import "strings"

func Repeat(c string) string {
	var rep strings.Builder

	for range 5 {
		rep.WriteString(c)
	}

	return rep.String()
}
