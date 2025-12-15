package sqli


// Detect tries minimal payloads and reports suspicion
func Detect(send func(string) (string, error), baseline string) (bool, string) {
for _, p := range DetectPayloads {
body, err := send(p)
if err != nil {
continue
}
if Anomalous(baseline, body) {
return true, p
}
}
return false, ""
}