package analyzer

import (
	"fmt"
	"io"
	"regexp"

	"github.com/farzadamr/ctf-web-tool/internal/httpclient"
)

func Analyze(target string) {
	resp, err := httpclient.Client.Get(target)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("Headers:")
	for k, v := range resp.Header {
		fmt.Println(" ", k, v)
	}

	re := regexp.MustCompile(`https?://[a-zA-Z0-9./_-]+`)
	links := re.FindAllString(string(body), -1)
	fmt.Println("Links found:")
	for _, l := range links {
		fmt.Println(" ", l)
	}
}
