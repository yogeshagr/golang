// Lets users create, read, update, and delete GitHub issues

package main

import (
	"net/http"
	"fmt"
	"bytes"
	"io/ioutil"
	"os"
	)

func main() {
	url := "https://api.github.com/repos/yogeshagr/golang/issues"
	var jsonStr = []byte(`{
		"title": "Found a bug",
		"body": "I'm having a problem with this",
		"labels": ["bug"],
		"assignees": ["yogeshagr"]
		}`)

  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_PASS"))

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("response status:", resp.Status)
	fmt.Println("response headers", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println("response body", string(body))
}
