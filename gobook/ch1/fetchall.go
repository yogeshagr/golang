// Fetchall fetches URLs in parallel and reports their times and sizes.

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(uri string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(uri)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	f, err := os.Create(url.QueryEscape(uri))
	if err != nil {
		ch <- fmt.Sprint(err)
	}

	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close()

	if closeError := f.Close(); err == nil {
		err = closeError
	}

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", uri, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, uri)
}
