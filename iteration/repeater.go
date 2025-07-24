package iteration

import "strings"

func Repeat(c string, times int) string {
	var rep strings.Builder

	for range times {
		rep.WriteString(c)
	}

	return rep.String()
}
