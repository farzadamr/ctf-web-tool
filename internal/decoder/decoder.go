package decoder

import (
	"encoding/base64"
	"encoding/hex"
	"strings"
)

func Decode(value string) string {
	if out, err := base64.StdEncoding.DecodeString(value); err == nil {
		return "[base64]: " + string(out)
	}
	if out, err := hex.DecodeString(value); err == nil {
		return "[hex]: " + string(out)
	}
	if strings.Count(value, ".") == 2 {
		return "Looks like JWT token → use: jwt -v <token>"
	}
	return ""
}
