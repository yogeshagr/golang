package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	fmt.Fprintln(c, "\t upper", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t lower", strings.ToLower(shout))
}

func handleFunc(c net.Conn) {
	var wg sync.WaitGroup // number of working goroutines
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second, &wg)
	}
	wg.Wait()
	if tcpconn, ok := c.(*net.TCPConn); ok {
		fmt.Println("server: closing the write end")
		tcpconn.CloseWrite()
	}
}

func main() {
	port := os.Args[1]
	address := fmt.Sprintf("%s:%s", "localhost", port)
	fmt.Println(address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleFunc(conn) // handle one connection at a time
	}
}
