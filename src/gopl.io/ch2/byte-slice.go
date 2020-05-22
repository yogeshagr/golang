package main

import "fmt"

func main() {
	var buf []byte
	buf[0] = '['
	fmt.Fprintf(&buf, "%d", 1)
	buf[2] = ']'
	fmt.Println(buf)
}

/*
This program will fail as follows, which is expected.
$ go build -o ./bin/ byte-slice.go
# command-line-arguments
./byte-slice.go:8:14: cannot use &buf (type *[]byte) as type io.Writer in argument to fmt.Fprintf:
				*[]byte does not implement io.Writer (missing Write method)
*/
