package decoder

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

func Decode(value string) {
	if out, err := base64.StdEncoding.DecodeString(value); err == nil {
		fmt.Println("[base64]:", string(out))
	}
	if out, err := hex.DecodeString(value); err == nil {
		fmt.Println("[hex]:", string(out))
	}
	if strings.Count(value, ".") == 2 {
		fmt.Println("Looks like JWT token → use: jwt -v <token>")
	}
}
