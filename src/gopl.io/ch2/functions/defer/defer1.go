/*
f(%d)
 3
f(%d)
 2
f(%d)
 1
defer 1
defer 2
defer 3
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.f(0x0)
        /Users/yogeshagrawal/work/self/golang/src/gopl.io/ch2/functions/defer/defer1.go:10 +0x1e4
main.f(0x1)
        /Users/yogeshagrawal/work/self/golang/src/gopl.io/ch2/functions/defer/defer1.go:12 +0x185
main.f(0x2)
        /Users/yogeshagrawal/work/self/golang/src/gopl.io/ch2/functions/defer/defer1.go:12 +0x185
main.f(0x3)
        /Users/yogeshagrawal/work/self/golang/src/gopl.io/ch2/functions/defer/defer1.go:12 +0x185
main.main()
        /Users/yogeshagrawal/work/self/golang/src/gopl.io/ch2/functions/defer/defer1.go:6 +0x2a
*/

package main

import "fmt"

func main() {
	f(3)
}

func f(x int) {
	fmt.Println("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
