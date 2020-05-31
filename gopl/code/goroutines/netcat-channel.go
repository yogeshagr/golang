/*
The netcat1 program copies input to the server in its main gorotine, so the
client program terminates as soon as the input stream closes, even if the
background goroutine is still working. To make the program wait for the
background goroutine to complete before exiting, we use a channel to synchronize
the two goroutines.
*/

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // Signal the main goroutine
	}()

	mustCopy(conn, os.Stdin)
	if tcpconn, ok := conn.(*net.TCPConn); ok {
		fmt.Println("closing the write end")
		tcpconn.CloseWrite()
	}
	fmt.Println("waiting for goroutine to finish")
	<-done // wait for background goroutine to finish
	fmt.Println("completed")
}

func mustCopy(dst io.Writer, src io.Reader) {
	fmt.Println("copying to server")
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
	fmt.Println("copying to server done")
}
