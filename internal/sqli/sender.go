package sqli

import (
	"io"
	"net/http"
	"net/url"
)

func BuildSender(target, param string) func(string) (string, error) {
	return func(payload string) (string, error) {
		u, err := url.Parse(target)
		if err != nil {
			return "", err
		}

		q := u.Query()
		q.Set(param, payload)
		u.RawQuery = q.Encode()

		resp, err := http.Get(u.String())
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		b, _ := io.ReadAll(resp.Body)
		return string(b), nil
	}
}
