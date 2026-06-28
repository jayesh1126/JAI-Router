# jai-router

> **Point your AI coding tool at it. Cheap calls go local. See exactly what you saved.**

A local-first proxy that sits between your AI coding tools (Claude Code, Cline, Aider, OpenCode) and the cloud. It routes background calls — titles, summaries, short Q&A — to a local Ollama model, and sends the real work to the cloud. One binary, no infrastructure, honest savings.

```bash
# Coming soon
brew install jai-router
jai-router init   # detects Ollama, asks for your API keys once
jai-router start  # OpenAI endpoint: localhost:8080/v1
                 # Anthropic endpoint: localhost:8080
                 # Dashboard: localhost:8080/dashboard
```

Then in Cline, Aider, or OpenCode:
```json
{ "openAiBaseUrl": "http://localhost:8080/v1" }
```

For Claude Code:
```bash
export ANTHROPIC_BASE_URL=http://localhost:8080
```

---

## How it works

Every AI coding session is a mix of calls: some need GPT-4 or Claude — tool use, file edits, complex reasoning. Many don't — commit message generation, quick summaries, short answers. Today, tools send everything to the cloud at the same price.

jai-router intercepts each request and routes it:

- **Requests with tool calls, vision, or long context** → cloud (correctness matters)
- **Short, tool-free requests** → local Ollama model (fast and free)
- **Streaming passthrough** — tokens flow straight through, no added latency

The dashboard shows your real savings: actual model used, baseline cost if it had gone to cloud, running total.

---

## Why not LiteLLM / Olla / TriageLLM?

Those tools solve infrastructure problems (multi-GPU homelab routing, cloud API aggregation). This solves a developer workflow problem: you're using Cline or Claude Code every day and paying full cloud price for every single call. jai-router is a single binary with no external dependencies, designed specifically for this use case, with per-tool setup docs and an honest savings counter you can screenshot.

---

## Status

**Early development.** Core passthrough working. Routing logic and dashboard in progress.

- [x] OpenAI-compatible endpoint
- [ ] Ollama provider + routing engine
- [ ] Metrics + SQLite persistence
- [ ] Dashboard
- [ ] Anthropic endpoint (Claude Code support)
- [ ] `jai-router wrap <tool>` one-command setup

Feedback, issues, and PRs welcome. See [SPEC.md](./SPEC.md) for the full architecture.

---

## License

MIT