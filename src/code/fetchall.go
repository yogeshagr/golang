package main

import (
	"fmt"
	"io"
	"strings"
	"net/http"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	start := time.Now()
	file, error := os.OpenFile("output.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)

	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	ch := make(chan string) // create a channel of strings
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://"){
			url = "http://" + url
		}
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:]{
		fmt.Fprintln(file, <-ch) // receive from channel ch
	}
	fmt.Fprintf(file, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
  resp, error := http.Get(url)
	if error != nil {
		ch <- fmt.Sprint(error) // send to channel ch
		return
	}

	nbytes, error := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if error != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, error)
		return
	}
	elapsedTime := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", elapsedTime, nbytes, url)
}
