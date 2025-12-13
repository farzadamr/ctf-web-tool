package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/farzadamr/ctf-web-tool/internal/analyzer"
	"github.com/farzadamr/ctf-web-tool/internal/decoder"
	"github.com/farzadamr/ctf-web-tool/internal/jwt"
	"github.com/farzadamr/ctf-web-tool/internal/static"
	"github.com/farzadamr/ctf-web-tool/internal/template"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ctf-web-tool <command> [args]")
		return
	}

	cmd := os.Args[1]
	switch cmd {
	case "analyze":
		url := flag.String("url", "", "Target URL")
		flag.CommandLine.Parse(os.Args[2:])
		analyzer.Analyze(*url)

	case "decode":
		value := flag.String("v", "", "Encoded token")
		flag.CommandLine.Parse(os.Args[2:])
		decoder.Decode(*value)

	case "template":
		url := flag.String("url", "", "Target URL")
		param := flag.String("param", "q", "Parameter name")
		flag.CommandLine.Parse(os.Args[2:])
		template.TestInjection(*url, *param)

	case "jwt":
		value := flag.String("v", "", "JWT token")
		flag.CommandLine.Parse(os.Args[2:])
		jwt.CheckJWT(*value)

	case "scan-static":
		url := flag.String("url", "", "Target URL")
		flag.CommandLine.Parse(os.Args[2:])
		static.Scan(*url)

	default:
		fmt.Println("Unknown command: ", cmd)
	}
}
