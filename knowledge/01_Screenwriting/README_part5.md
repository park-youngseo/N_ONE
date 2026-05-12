
```bash
# 1. Download the script first — review it before running
curl -sSfL https://raw.githubusercontent.com/Insajin/autopus-adk/main/install.sh -o install.sh
less install.sh          # Read what it does
sh install.sh            # Run only after review
```

**Or verify manually:**

```bash
# Download binary + checksums separately
VERSION=$(curl -s https://api.github.com/repos/Insajin/autopus-adk/releases/latest | grep tag_name | sed 's/.*"v\(.*\)".*/\1/')
curl -LO "https://github.com/Insajin/autopus-adk/releases/download/v${VERSION}/autopus-adk_${VERSION}_$(uname -s | tr A-Z a-z)_$(uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/').tar.gz"
curl -LO "https://github.com/Insajin/autopus-adk/releases/download/v${VERSION}/checksums.txt"

# Verify SHA256
shasum -a 256 -c checksums.txt --ignore-missing
```

`auto update --self` also verifies SHA256 checksums before replacing the binary.

### What We Don't Do

- No telemetry or analytics collection
- No network calls except explicit commands (`orchestra`, `search`, `update --self`)
- No access to your AI provider API keys — Autopus orchestrates CLI tools, not API calls

---

## 🤝 Contributing

Autopus-ADK is open source under the MIT license. PRs welcome!

```bash
make test       # Run tests with race detection
make lint       # Run go vet
make coverage   # Generate coverage report
```

---

<div align="center">

**🐙 Autopus** — Of the agents. By the agents. For the agents.

</div>