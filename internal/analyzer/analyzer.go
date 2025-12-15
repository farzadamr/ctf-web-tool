package analyzer

import (
	"io"
	"regexp"
	"strings"

	"github.com/farzadamr/ctf-web-tool/internal/httpclient"
)

func Analyze(target string) string {
	resp, err := httpclient.Client.Get(target)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}

	var out strings.Builder

	// Status
	out.WriteString("Status: " + resp.Status + "\n")

	// Headers
	out.WriteString("Headers:\n")
	for k, v := range resp.Header {
		out.WriteString("  " + k + ": " + strings.Join(v, ", ") + "\n")
	}

	// Links
	re := regexp.MustCompile(`https?://[a-zA-Z0-9./_-]+`)
	links := re.FindAllString(string(body), -1)
	if len(links) == 0 {
		out.WriteString("No links found\n")
	} else {
		out.WriteString("Links found:\n")
		for _, l := range links {
			out.WriteString("  " + l + "\n")
		}
	}

	return out.String()
}
