package static

import (
	"io"
	"regexp"
	"strings"

	"github.com/farzadamr/ctf-web-tool/internal/httpclient"
)

func Scan(url string) string {
	resp, err := httpclient.Client.Get(url)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}

	re := regexp.MustCompile(`(?i)<!--.*?-->`)
	comments := re.FindAllString(string(body), -1)

	if len(comments) == 0 {
		return "No HTML comments found"
	}

	var out strings.Builder
	out.WriteString("Comments:\n")
	for _, c := range comments {
		out.WriteString("  " + c + "\n")
	}

	return out.String()
}
