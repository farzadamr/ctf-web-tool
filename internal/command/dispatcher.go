package command

import (
	"flag"
	"os"

	"github.com/farzadamr/ctf-web-tool/internal/analyzer"
	"github.com/farzadamr/ctf-web-tool/internal/decoder"
	"github.com/farzadamr/ctf-web-tool/internal/jwt"
	"github.com/farzadamr/ctf-web-tool/internal/mapper"
	sqli "github.com/farzadamr/ctf-web-tool/internal/sqli"
	"github.com/farzadamr/ctf-web-tool/internal/static"
	"github.com/farzadamr/ctf-web-tool/internal/template"
)

func Run(args []string) string {
	if len(args) < 1 {
		return "usage: ctf-web-tool <command> [args]"
	}

	switch args[0] {

	case "sqli":
		fs := flag.NewFlagSet("sqli", flag.ContinueOnError)
		url := fs.String("url", "", "Target URL")
		param := fs.String("param", "id", "Parameter name")
		_ = fs.Parse(args[1:])

		if *url == "" {
			return "Usage: ctf-web-tool sqli --url <target> --param <param>"
		}

		send := sqli.BuildSender(*url, *param)

		baseline, err := send("1")
		if err != nil {
			return err.Error()
		}

		ok, payload := sqli.Detect(send, baseline)
		if !ok {
			return "[-] No SQL Injection detected"
		}

		out := "[+] Possible SQL Injection detected\n"
		out += "    Payload: " + payload + "\n"

		found, flagVal := sqli.ExtractFlag(send)
		if found {
			out += "[✔] FLAG FOUND: " + flagVal
		} else {
			out += "[!] SQLi confirmed but flag not found"
		}

		return out
	case "map":
		fs := flag.NewFlagSet("map", flag.ContinueOnError)
		targetURL := fs.String("url", "", "Target URL")
		_ = fs.Parse(args[1:])

		if *targetURL == "" {
			return "Usage: ctf-web-tool map --url <target>"
		}

		return mapper.MapTargetString(*targetURL)

	case "analyze":
		url := flag.String("url", "", "Target URL")
		flag.CommandLine.Parse(os.Args[2:])
		return analyzer.Analyze(*url)

	case "decode":
		value := flag.String("v", "", "Encoded token")
		flag.CommandLine.Parse(os.Args[2:])
		return decoder.Decode(*value)

	case "template":
		url := flag.String("url", "", "Target URL")
		param := flag.String("param", "q", "Parameter name")
		flag.CommandLine.Parse(os.Args[2:])
		return template.TestInjection(*url, *param)

	case "jwt":
		value := flag.String("v", "", "JWT token")
		flag.CommandLine.Parse(os.Args[2:])
		return jwt.CheckJWT(*value)

	case "scan-static":
		url := flag.String("url", "", "Target URL")
		flag.CommandLine.Parse(os.Args[2:])
		return static.Scan(*url)
	default:
		return "Unknown command: " + args[0]
	}
}
