package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	pre := len(s) % 3
	if pre == 0 {
		pre = 3
	}
	buf.WriteString(s[:pre])
	for i := pre; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

func main() {
	fmt.Printf("%v\n", comma("1234556"))
}
