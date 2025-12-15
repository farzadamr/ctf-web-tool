package sqli


import "regexp"


var FlagRegexes = []*regexp.Regexp{
regexp.MustCompile(`flag\{[^}]+}`),
regexp.MustCompile(`FLAG\{[^}]+}`),
regexp.MustCompile(`ctf\{[^}]+}`),
regexp.MustCompile(`CTF\{[^}]+}`),
}