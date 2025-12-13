# ctf-web-tool

A lightweight, CTF-oriented CLI tool written in Go for **quickly analyzing Web challenge attack surfaces** (easy → medium), without fuzzing or brute-force.

Designed for **rate-limited CTF environments** where speed and signal matter.

---

## ✨ Features

- 🔍 **Attack Surface Mapping**
  - Extracts endpoints from HTML & JavaScript
  - Ranks endpoints by risk (HIGH / MEDIUM / LOW)
  - Highlights sensitive keywords (`admin`, `debug`, `upload`, etc.)

- 🍪 **Token & Cookie Decoding**
  - Base64 / Base64URL
  - Hex
  - JWT (header & payload inspection)

- 🧠 **CTF-Focused Heuristics**
  - Minimal HTTP requests
  - No fuzzing, no brute-force
  - Hint-driven output

---

## 🚀 Usage

```bash
ctf-tool map --url https://target
ctf-tool decode -v <token>
ctf-tool jwt -v <jwt>
ctf-tool template --url https://target --param name
ctf-tool scan-static --url https://target
ctf-tool analyze --url https://target
