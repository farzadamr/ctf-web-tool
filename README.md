# CTF Web Tool
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

CTF Web Tool is a command-line utility designed for web security testing and CTF challenges.  
It provides multiple modules for reconnaissance, token analysis, JWT inspection, template injection testing, static analysis, and SQL injection detection.

This tool is intended for educational purposes and authorized security testing only.

---

## 🚀 Features

- 🌐 Web endpoint mapping
- 🔐 Token and encoded value decoding
- 🧾 JWT analysis and inspection
- 🧪 Template injection testing
- 🕵️ Static web security scanning
- 📊 Target behavior analysis
- 💉 Basic SQL injection parameter testing
- 🖥 Windows service support for background execution

---

## 📦 Installation 

```bash
build -o ctf-web-tool_service.exe ./cmd/service
build -o ctf-web-tool.exe ./cmd/cli
```

```bash
./ctf-web-service.exe install
```
## 🛠 Usage

```bash
ctf-tool map --url https://target
ctf-tool decode -v <token>
ctf-tool jwt -v <jwt>
ctf-tool template --url https://target --param name
ctf-tool scan-static --url https://target
ctf-tool analyze --url https://target
ctf-tool sqli --url <target> --param <params>
```
## 📄 License

This project is licensed under the MIT License.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files, to deal in the Software
without restriction, including without limitation the rights to use, copy,
modify, merge, publish, distribute, sublicense, and/or sell copies of the Software.

The software is provided "as is", without warranty of any kind.

See the LICENSE file for more information.
