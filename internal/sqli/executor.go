package sqli


// ExtractFlag attempts to find a flag in responses
func ExtractFlag(send func(string) (string, error)) (bool, string) {
for _, p := range ExtractPayloads {
body, err := send(p)
if err != nil {
continue
}
for _, re := range FlagRegexes {
if m := re.FindString(body); m != "" {
return true, m
}
}
}
return false, ""
}