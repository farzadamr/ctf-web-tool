package static

import (
	"fmt"
	"io"
	"regexp"

	"github.com/farzadamr/ctf-web-tool/internal/httpclient"
)

func Scan(url string) {
	resp, err := httpclient.Client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	re := regexp.MustCompile(`(?i)<!--.*?-->`)
	comments := re.FindAllString(string(body), -1)
	fmt.Println("Comments:")
	for _, c := range comments {
		fmt.Println(" ", c)
	}
}
