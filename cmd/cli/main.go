package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ctf-web-tool <command>")
		return
	}


	body := bytes.NewBufferString(
		strings.Join(os.Args[1:], " "),
	)

	resp, err := http.Post(
		"http://127.0.0.1:7878/run",
		"text/plain",
		body,
	)

	if err != nil {
		fmt.Println("engine not running")
		return
	}

	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	fmt.Println(string(data))
}
