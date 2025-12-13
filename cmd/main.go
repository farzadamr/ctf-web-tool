package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/farzadamr/ctf-web-tool/internal/analyzer"
	"github.com/farzadamr/ctf-web-tool/internal/decoder"
	"github.com/farzadamr/ctf-web-tool/internal/jwt"
	"github.com/farzadamr/ctf-web-tool/internal/mapper"
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
	case "map":
		mapCmd := flag.NewFlagSet("map", flag.ExitOnError)
		targetURL := mapCmd.String("url", "", "Target URL")
		_ = mapCmd.Parse(os.Args[2:])

		if *targetURL == "" {
			fmt.Println("Usage: ctf-tool map --url <target>")
			os.Exit(1)
		}

		endpoints, err := mapper.MapTarget(*targetURL)
		if err != nil {
			log.Fatal(err)
		}

		for _, e := range endpoints {
			level := "LOW"
			if e.Score >= 7 {
				level = "HIGH"
			} else if e.Score >= 4 {
				level = "MEDIUM"
			}

			fmt.Printf("[%s] %s (score=%d)\n", level, e.Path, e.Score)
			for _, h := range e.Hints {
				fmt.Println("  -", h)
			}
		}
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
