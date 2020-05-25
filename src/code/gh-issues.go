// Lets users create, read, update, and delete GitHub issues

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const APIURL = "https://api.github.com"

type issue struct {
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Labels    []string `json:"labels"`
	Assignees []string `json:"assignees"`
}

func createIssue(owner string, repo string) {
	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues"}, "/")
	var myIssue issue
	myIssue.Title = "Found a bug"
	myIssue.Body = "I'm having a problem with this"
	myIssue.Labels = []string{"bug"}
	myIssue.Assignees = []string{"yogeshagr"}
	myJson, err := json.Marshal(myIssue)

	/*	var jsonStr := []byte(`{
		"title": "Found a bug",
		"body": "I'm having a problem with this",
		"labels": ["bug"],
		"assignees": ["yogeshagr"]
		}`) */

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(myJson))
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

var usage string = `usage:
[create|read] OWNER REPO
`

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	owner, repo := args[0], args[1]
	switch cmd {
	case "create":
		createIssue(owner, repo)
	}
}
