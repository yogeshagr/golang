package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, error := http.Get(url)
		if error != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", error)
			os.Exit(1)
		}
		b, error := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if error != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, error)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
