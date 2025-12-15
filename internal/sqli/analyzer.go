package sqli


import "strings"


// Simple anomaly analyzer
func Anomalous(a, b string) bool {
if len(a) == 0 || len(b) == 0 {
return false
}
if len(a) != len(b) {
return true
}
keywords := []string{"sql", "syntax", "mysql", "sqlite", "psql"}
al := strings.ToLower(a)
bl := strings.ToLower(b)
for _, k := range keywords {
if strings.Contains(bl, k) && !strings.Contains(al, k) {
return true
}
}
return false
}