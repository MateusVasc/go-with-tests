package iteration

func Repeat(c string) string {
	temp := c

	for range 5 {
		c += temp
	}

	return c
}
