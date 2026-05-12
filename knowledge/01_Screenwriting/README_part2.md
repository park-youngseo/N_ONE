```
feat(auth): add OAuth2 provider abstraction

Why: Need Google + GitHub support, extensible for future providers
Decision: Interface-based abstraction over direct SDK usage
Alternatives: Direct SDK calls (rejected: too coupled)
Ref: SPEC-AUTH-001

🐙 Autopus <noreply@autopus.co>
```

9 structured trailers. Query with `auto lore query "why interface?"`. Stale decisions auto-detected after 90 days.

### 🧪 Autonomous Experiment Loop

Let AI iterate autonomously — measure, keep or discard, repeat.

```bash
/auto experiment --metric "go test -bench=BenchmarkProcess" --direction lower --max-iter 5
```

```
🐙 Experiment ───────────────────────
  Iter 1: baseline  │ 1200 ns/op
  Iter 2: optimize  │  850 ns/op  ✓ keep (29% improvement)
  Iter 3: refactor  │  900 ns/op  ✗ discard (regression)
  Iter 4: cache     │  620 ns/op  ✓ keep (27% improvement)
  ─────────────────────────────────────
  Result: 1200 → 620 ns/op (48% improvement)
```

Built-in **circuit breaker** prevents runaway iterations. **Simplicity scoring** penalizes over-complex solutions. Each iteration is a git commit — easy to review or revert.

> ⚠️ **Status: Experimental** — CLI commands (`auto experiment`) are available but skill-level integration is in progress. Core iteration loop works; full pipeline integration is coming.

### 🧠 Pipeline That Learns From Failures

Autopus pipelines don't just fail — they **remember why** and prevent the same mistake next time.

```
Gate 2 FAIL: golangci-lint — unused variable in pkg/auth/
→ Auto-recorded to .autopus/learnings/pipeline.jsonl
→ Next /auto go: learning injected into executor prompt
→ Same mistake never repeated
```

Every pipeline failure is captured as a structured learning entry. On the next run, relevant learnings are automatically injected into agent prompts — giving your pipeline **institutional memory** across sessions.

### 🏥 Post-Deploy Health Check

Deploy first, verify immediately. `canary` runs build verification, E2E tests, and browser health checks against your live deployment.

```bash
/auto canary                          # Build + E2E + browser auto-verification
/auto canary --url https://myapp.com  # Target a specific deployment URL
/auto canary --watch 5m               # Repeat every 5 minutes
/auto canary --compare                # Compare against previous canary report
```

Generates `canary.md` with full diagnostics — build status, test results, accessibility scores, and screenshot diffs.

### 🔀 Smart Model Routing

Not every task needs Opus. Autopus analyzes message complexity and routes to the right model automatically.

```
Simple query     → Haiku  (fast, cheap)
Code review      → Sonnet (balanced)
Architecture     → Opus   (deep reasoning)
```

No configuration needed — the router evaluates token count, code complexity, and domain signals to pick the optimal model. Override anytime with `--quality ultra`.

### 🔌 Provider Connection Wizard

Setting up AI providers shouldn't require reading docs. `auto connect` walks you through a 3-step guided setup.

```bash
auto connect    # Interactive wizard: detect → configure → verify
```

Detects installed CLI tools, validates API keys, tests connectivity, and writes provider config — all in one command.

### 🤖 ADK Worker — Local Agent Execution

ADK Worker runs A2A + MCP hybrid tasks locally with browser login, JWT refresh, and direct platform connectivity.
No separate bridge daemon or worker API key exchange is required for the default production path.

What it is for:
- Connecting a local workspace to the Autopus platform worker loop
- Receiving platform-dispatched tasks and executing them with local tools
- Reusing the same security, budget, and audit rails as the main harness

What to do today:
- If you're here for `auto init`, Codex `@auto ...`, or OpenCode `/auto ...`, you can ignore Worker for now
- `auto worker ...` is an optional advanced surface that is still being rolled out and documented

### 💰 Iteration Budget Management

Workers don't run forever. Each executor gets a tool-call budget — preventing runaway agents while ensuring enough room to complete complex tasks.

### 📦 Context Compression

As pipelines progress through phases, earlier context gets compressed automatically — keeping agent prompts focused and within token limits without losing critical information.

### 🔄 Pipeline That Never Dies

Crash mid-pipeline? Resume exactly where you left off.

```bash
/auto go SPEC-AUTH-001 --continue    # Resume from last checkpoint
```

YAML-based checkpoints save pipeline state after every phase. Stale detection prevents resuming outdated sessions. Combined with `--auto --loop`, you get a **fully resilient autonomous pipeline.**

### 🧪 E2E Scenarios from Your Code

Auto-generate and execute E2E test scenarios — no manual test writing needed.

```bash
auto test run                    # Run all scenarios
auto test run -s init --verbose  # Run a specific scenario
```

Autopus analyzes your codebase (Cobra commands, API routes, frontend pages) and generates typed scenarios with **verification primitives** (`exit_code`, `stdout_contains`, `status_code`, `json_path`, etc.). Incremental sync keeps scenarios up-to-date as code evolves.

### 🌐 Browser Automation — AI Agents That See and Click

AI agents can directly interact with web pages — open URLs, read accessibility trees, click elements, fill forms, and capture screenshots.

```bash
/auto browse --url https://example.com/settings
```

```
- @e1 heading "AI Settings"
- @e2 button "Provider Mode"
- @e3 switch "Auto Fallback" [checked]
- @e7 button "Save"
```

Terminal-aware: automatically selects `cmux browser` (in cmux) or `agent-browser` (fallback). Snapshot → Act → Verify loop — agents see the page as an accessibility tree and interact by reference.

### 📺 Live Agent Dashboard

In `--team` mode, each team member gets its own terminal pane with real-time log streaming.

```
┌─ lead ──────────┬─ builder-1 ───────┐
│ Phase 1: Plan   │ T1: auth.go       │
│ 5 tasks created │ implementing...   │
├─ tester ────────┼─ guardian ────────┤
│ scaffold: 12    │ waiting...        │
│ RED state ✓     │                   │
└─────────────────┴───────────────────┘
```

Works in cmux and tmux. Plain terminals degrade gracefully to log-only output.

### 📚 Auto-Documentation with Context7

Before implementation, Autopus fetches latest library docs automatically — so agents never work with stale API knowledge.

```
Phase 1.8: Doc Fetch
  → Detected: cobra v1.9, testify v1.11
  → Fetched: 2 libraries (6000 tokens)
  → Injected into executor + tester prompts
```

Context7 MCP → WebSearch fallback → skip (never blocks pipeline). Adaptive token budget: 1 lib → 5000 tokens, 5 libs → 2000 tokens each.

### 🔌 Hook-Based Result Collection

Instead of scraping terminal output, Autopus uses each provider's native hook system to collect structured JSON results.

| Provider | Hook Type | How |
|----------|-----------|-----|
| Claude Code | Stop hook | Extracts `last_assistant_message` |
| Gemini CLI | AfterAgent hook | Extracts `prompt_response` |
| OpenCode | Plugin | Extracts `text` field |

Fallback: providers without hooks use ReadScreen + idle detection (SPEC-ORCH-006).

### 🔧 More Power Tools

| Feature | Command | What It Does |
|---------|---------|-------------|
| **Reaction Engine** | `auto react check/apply` | Detects CI failures, analyzes logs, generates fix reports automatically |
| **Meta-Agent Builder** | `auto agent create` / `auto skill create` | Scaffold custom agents and skills from patterns |
| **Hard Gate** | `auto check --gate` | Enforce mandatory pipeline gates (mandatory/advisory modes) |
| **Self-Update** | `auto update --self` | Atomic binary update — GitHub Releases check + SHA256 verification |
| **Cost Tracking** | `auto telemetry cost` | Token-based pipeline cost estimation per model |
| **Issue Reporter** | `auto issue report` | Auto-collect error context, sanitize secrets, create GitHub issues |
| **Signature Map** | `auto setup` | Extract exported API signatures (Go + TypeScript) via AST analysis |
| **Test Runner Detection** | `auto init` | Auto-detect jest, vitest, pytest, cargo test frameworks |

### 🌐 One Config, Four Platforms

```bash
auto init   # auto-detects supported installed AI coding CLIs
```

One `autopus.yaml` generates **native configuration** for every detected supported platform.

| Platform | What Gets Generated |
|----------|-------------------|
| **Claude Code** | `.claude/rules/`, `.claude/skills/`, `.claude/agents/`, `CLAUDE.md` |
| **Codex** | `.codex/`, `.agents/skills/`, `.agents/plugins/marketplace.json`, `.autopus/plugins/auto/`, `AGENTS.md` |
| **Gemini CLI** | `.gemini/`, `GEMINI.md` |
| **OpenCode** | `.opencode/rules/`, `.opencode/agents/`, `.opencode/commands/`, `.opencode/plugins/`, `.agents/skills/`, `AGENTS.md`, `opencode.json` |
Same 16 agents. Same 40 skills. Same rules. **Every platform.**

Codex note:
- Use `$auto plan ...`, `$auto go ...`, `$auto idea ...` immediately after `auto init` or `auto update`
- Install the generated local plugin from the marketplace entry in `.agents/plugins/marketplace.json` (`.autopus/plugins/auto`) to unlock the friendlier `@auto ...` syntax

OpenCode note:
- `/auto ...` and direct aliases like `/auto-plan ...` are generated under `.opencode/commands/`
- Native rule/agent/plugin files live under `.opencode/`, while reusable skills are published under `.agents/skills/`
- Helper workflows like `/auto status`, `/auto map`, `/auto why`, `/auto verify`, `/auto secure`, `/auto test`, `/auto dev`, and `/auto doctor` are generated as OpenCode-native command wrappers
- `opencode.json` now registers the managed hook plugin automatically, so `.opencode/plugins/autopus-hooks.js` is live immediately after `auto init` or `auto update`

### Codex vs OpenCode

| Topic | Codex | OpenCode |
|-------|-------|----------|
| Primary command syntax | `@auto <subcommand> ...` | `/auto <subcommand> ...` |
| Works immediately after `auto init` | `$auto ...` repo-skill fallback | `/auto ...` and `/auto-<subcommand> ...` wrappers |
| Extra install step | Yes. Install the generated local plugin from `.agents/plugins/marketplace.json` to enable `@auto ...` | No extra router install step. `opencode.json` wires the managed plugin automatically |
| Generated surface | `.codex/`, `.agents/skills/`, `.agents/plugins/marketplace.json`, `.autopus/plugins/auto/`, `AGENTS.md` | `.opencode/commands/`, `.opencode/agents/`, `.opencode/rules/`, `.opencode/plugins/`, `.agents/skills/`, `AGENTS.md`, `opencode.json` |
| What works well today | Core `auto` workflows, repo skills, local plugin-based `@auto` routing | Core `auto` workflows, native command wrappers, managed hook plugin wiring |
| Current boundary | `@auto ...` depends on local plugin installation; without it, use `$auto ...` | Current parity target is the core workflow surface. Claude-style native settings/statusline breadth is not claimed |
| Worker surface | Optional for now. Ignore unless you specifically need platform-connected worker execution | Optional for now. Ignore unless you specifically need platform-connected worker execution |

---

## 🚀 Quick Start Guide

Get from zero to your first AI-powered feature in under 5 minutes.

### Step 1 · Install (one line)

> **Paste this command into your AI coding agent's chat** (Claude Code, Codex, OpenCode, etc.) — the agent will run it for you. Or run it directly in your terminal.

```bash
# macOS / Linux — installs the binary and checks required tools
cd your-project    # go to your project folder (e.g., cd ~/my-app)
curl -sSfL https://raw.githubusercontent.com/Insajin/autopus-adk/main/install.sh | sh

# Windows (CMD or PowerShell)
cd your-project
powershell -c "irm https://raw.githubusercontent.com/Insajin/autopus-adk/main/install.ps1 | iex"
```

**That's it.** The installer installs the `auto` CLI plus an `autopus` alias, checks required tools, skips anything already present, and auto-installs missing essentials like `git` and GitHub CLI. It does **not** run `auto init` for you.

Platform command syntax:
- Codex: install the generated local plugin, then use `@auto ...`; until then, use `$auto ...`
- OpenCode: use `/auto ...` or `/auto-<subcommand> ...`
- Claude Code / Gemini CLI: use `/auto ...`

> Note: If you run the Windows installer from Git Bash via `powershell -c ...`, restart Git Bash after install so it reloads the updated user `PATH`. The installer prints the exact install directory and a one-line `export PATH=...` fallback for that case.

<details>
<summary>Other install methods</summary>

```bash
# Homebrew (macOS)
brew install insajin/tap/autopus-adk

# go install (requires Go 1.26+)
go install github.com/Insajin/autopus-adk/cmd/auto@latest

# Build from source
git clone https://github.com/Insajin/autopus-adk.git
cd autopus-adk && make build && make install

# After manual install, initialize:
cd your-project && auto init
```

</details>

<details>
<summary>Installer options (environment variables)</summary>

| Variable | Default | Description |
|----------|---------|-------------|
| `INSTALL_DIR` | `/usr/local/bin` | Binary install path |
| `VERSION` | latest | Specific version to install |

</details>

After install, the script explains these commands:

- `auto init`: initialize the current project and generate `autopus.yaml` plus platform files
- `auto update --self`: update the `auto` CLI binary itself
- `auto update`: refresh rules, skills, agents, and other generated harness files in your project
