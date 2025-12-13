package jwt

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func b64(s string) string {
	out, _ := base64.RawURLEncoding.DecodeString(s)
	return string(out)
}

func CheckJWT(t string) {
	parts := strings.Split(t, ".")
	if len(parts) != 3 {
		fmt.Println("Not JWT")
		return
	}

	fmt.Println("Header:", b64(parts[0]))
	fmt.Println("Payload:", b64(parts[1]))
	fmt.Println("Signature:", parts[2])
}
