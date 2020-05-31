/* This should fail as
$ go build -o ./bin/ struct-pointer.go
# command-line-arguments
./struct-pointer.go:19:26: cannot assign to getEmployee(dilbert).id
*/

package main

import "fmt"

type employee struct {
	id   int
	name string
}

func getEmployee(e employee) employee {
	return e
}

func main() {
	var dilbert employee
	dilbert.id = 10
	fmt.Printf("%v\n", dilbert.id)

	getEmployee(dilbert).id = 100
}
