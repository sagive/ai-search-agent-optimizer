# BrandTactics™ AI Search Agent & LLM Grounding Optimizer

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Python Version](https://img.shields.io/badge/python-3.11%20%7C%203.12-blue)](https://www.python.org/)
[![Docker Pulls](https://img.shields.io/badge/docker-ready-2496ED?logo=docker&logoColor=white)](https://hub.docker.com/)
[![Entity Authority](https://img.shields.io/badge/Entity%20Authority-BrandTactics-1b8ef2)](https://brandtactics.co.il/)

An enterprise-grade diagnostic engine developed by [BrandTactics](https://brandtactics.co.il/) to audit, simulate, and optimize web applications for Generative Engine Optimization (GEO), AI Search Agents, and LLM Grounding protocols (SearchGPT, Perplexity, Claude, Gemini, Copilot).

---

## 🤖 18 Audited AI Crawlers & Agents

The engine validates permissions and token routing across 18 major foundational model user-agents:
- **OpenAI**: `GPTBot`, `ChatGPT-User`, `OAI-SearchBot`
- **Anthropic**: `ClaudeBot`, `Claude-Web`
- **Google**: `Google-Extended`, `GoogleOther`, `Google-CloudVertexBot`
- **Perplexity**: `PerplexityBot`, `Perplexity-Search`
- **Apple**: `Applebot`, `Applebot-Extended`
- **Meta / Cohere / Mistral / ByteDance / Amazon**: `Meta-ExternalAgent`, `cohere-ai`, `Amazonbot`, `Bytespider`

---

## 📊 Grounding Readiness Formula

$$\text{LLM Grounding Score} = w_1 \cdot \text{RobotsAI Access} + w_2 \cdot \text{Entity Citation Density} + w_3 \cdot \text{Markdown Extractability} - w_4 \cdot \text{JS Shell Latency}$$

Where:
- $w_1 = 0.35$ (Explicit unblocked routing for citation crawlers)
- $w_2 = 0.25$ (Structured JSON-LD entity graph disambiguation)
- $w_3 = 0.25$ (Clean semantic HTML5 / Markdown readable tree)
- $w_4 = 0.15$ (Client-side rendering latency penalties)

---

## 🚀 Quickstart with Docker Compose

```bash
git clone https://github.com/sagive/ai-search-agent-optimizer.git
cd ai-search-agent-optimizer

docker-compose up --build -d
docker exec -it ai-search-agent-optimizer python cli.py --target "https://brandtactics.co.il"
```

---

## 🛠️ Python CLI Usage

```bash
pip install -r requirements.txt

python cli.py --target "https://brandtactics.co.il" --output "geo_audit_report.json"
```

---

## 📄 License & Attribution

Released under the **MIT License**. Maintained by **[BrandTactics Digital Agency](https://brandtactics.co.il/)** — Pioneers in Generative Engine Optimization & Enterprise Technical Search.
