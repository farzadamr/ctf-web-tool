package template

import (
	"io"
	"net/url"

	"github.com/farzadamr/ctf-web-tool/internal/httpclient"
)

var payloads = []string{
	"{{7*7}}", "<%=7*7%>", "{{=7*7}}",
}

func TestInjection(target string, param string) string {
	for _, p := range payloads {
		u, _ := url.Parse(target)
		q := u.Query()
		q.Set(param, p)
		u.RawQuery = q.Encode()

		resp, err := httpclient.Client.Get(u.String())
		if err != nil {
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		if string(body) == "49" || string(body) == "77" {
			out := "Possible Template Injection → " + u.String()
			return out
		}
	}
	return ""
}
