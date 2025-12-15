package jwt

import (
	"encoding/base64"
	"strings"
)

func b64(s string) string {
	out, _ := base64.RawURLEncoding.DecodeString(s)
	return string(out)
}

func CheckJWT(t string) string {
	parts := strings.Split(t, ".")
	if len(parts) != 3 {
		return "Not JWT"
	}

	out := "Header: " + b64(parts[0]) + "\n"
	out += "Payload: " + b64(parts[1]) + "\n"
	out += "Signature: " + parts[2]

	return out
}
