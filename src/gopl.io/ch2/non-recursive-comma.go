package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer

	start := 0
	end := len(s)
	if s[0] == '-' || s[0] == '+' {
		buf.WriteByte('-')
		start = 1
	}

	float := false
	dot := -1
	if dot = strings.Index(s, "."); dot > 0 {
		float = true
		end = dot
	}
	pre := (end - start) % 3
	if pre == 0 {
		pre = 3
	}
	buf.WriteString(s[start : start+pre])
	for i := start + pre; i < end; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	if float {
		buf.WriteString(s[dot:])
	}
	return buf.String()
}

func main() {
	fmt.Printf("%v\n", comma("1234556.9791872"))
}
