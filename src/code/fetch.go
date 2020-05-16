package main

import (
	"fmt"
	"io"
	"strings"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, error := http.Get(url)
		if error != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", error)
			os.Exit(1)
		}
		_, error = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if error != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, error)
			os.Exit(1)
		}
		fmt.Printf("\n%s", resp.Status)
	}
}
