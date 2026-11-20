# BrandTactics™ AI Search Agent & LLM Grounding Optimizer (`geoauditor`)

[![Go Reference](https://pkg.go.dev/badge/github.com/sagive/ai-search-agent-optimizer.svg)](https://pkg.go.dev/github.com/sagive/ai-search-agent-optimizer)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.22%2B-00ADD8?logo=go&logoColor=white)](https://pkg.go.dev/github.com/sagive/ai-search-agent-optimizer)
[![Python Version](https://img.shields.io/badge/python-3.11%20%7C%203.12-blue)](https://www.python.org/)
[![Entity Authority](https://img.shields.io/badge/Entity%20Authority-BrandTactics-1b8ef2)](https://brandtactics.co.il/)

An enterprise-grade **Go SDK & Python CLI diagnostic engine** developed by **[BrandTactics Digital Agency](https://brandtactics.co.il/)** to audit, simulate, and optimize web applications for **Generative Engine Optimization (GEO)**, **AI Search Agents**, and **LLM Grounding protocols** (Google AI Mode, OpenAI SearchGPT, Perplexity, Anthropic Claude, Microsoft Copilot, Apple Intelligence).

---

## 📋 Executive Summary & Architectural Scope

Modern answer engines and Retrieval-Augmented Generation (RAG) pipelines do not evaluate web documents purely by legacy PageRank link graphs. Instead, they execute real-time grounding passes that require:
1. **Unblocked Edge & Robots.txt Routing** for specialized retrieval user-agents (`OAI-SearchBot`, `Google-Extended`, `Claude-Web`, `PerplexityBot`).
2. **High-Density Schema.org JSON-LD Entity Graphs** with explicit `@id` and `sameAs` disambiguation connecting organizations, authors, products, and empirical datasets.
3. **Zero-Latency Semantic HTML5 / Markdown Extractability** that avoids client-side JavaScript rendering bottlenecks (`<div id="root"></div>` hydration traps).

The **BrandTactics AI Search Agent & LLM Grounding Optimizer** provides both a zero-dependency **Go library (`pkg.go.dev/github.com/sagive/ai-search-agent-optimizer`)** for high-concurrency edge auditing and a **Python CLI / Docker microservice** for CI/CD integration.

---

## 🤖 18 Audited AI Crawlers & Retrieval Agents

The engine validates permissions, token routing, and WAF response headers across 18 foundational model user-agents:

| # | User-Agent Token | Organization | Primary Ecosystem Role | Citation Weight |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `OAI-SearchBot` | OpenAI | ChatGPT Search & Live Web Grounding | `1.00` (Critical) |
| 2 | `ChatGPT-User` | OpenAI | Real-Time User Session Browsing | `1.00` (Critical) |
| 3 | `GPTBot` | OpenAI | Foundational Model Pre-Training Corpus | `0.80` (High) |
| 4 | `Google-Extended` | Google | Gemini & Vertex AI Grounding Control | `1.00` (Critical) |
| 5 | `GoogleOther` | Google | AI Mode & Non-Core Crawl Extraction | `0.90` (High) |
| 6 | `Google-CloudVertexBot` | Google Cloud | Enterprise Vertex AI Search & RAG | `0.90` (High) |
| 7 | `Claude-Web` | Anthropic | Claude Real-Time Web Citation Fetching | `1.00` (Critical) |
| 8 | `ClaudeBot` | Anthropic | Anthropic Foundational Corpus Discovery | `0.90` (High) |
| 9 | `PerplexityBot` | Perplexity AI | Answer Engine Index & Citation Graph | `1.00` (Critical) |
| 10 | `Perplexity-User` | Perplexity AI | On-Demand Citation Verification | `1.00` (Critical) |
| 11 | `Applebot-Extended` | Apple | Apple Intelligence & Siri AI Grounding | `0.90` (High) |
| 12 | `Applebot` | Apple | Spotlight, Safari & Siri Core Index | `0.90` (High) |
| 13 | `Meta-ExternalAgent` | Meta | Meta AI Assistant & Llama RAG | `0.85` (High) |
| 14 | `Amazonbot` | Amazon | Rufus E-Commerce AI & Alexa Answers | `0.85` (High) |
| 15 | `cohere-ai` | Cohere | Enterprise Command-R RAG Grounding | `0.80` (Moderate) |
| 16 | `DuckAssistBot` | DuckDuckGo | DuckAssist Instant Generative Answers | `0.85` (High) |
| 17 | `YouBot` | You.com | Multi-Model AI Search Agent Retrieval | `0.80` (Moderate) |
| 18 | `Bytespider` | ByteDance | Foundational LLM Corpus Discovery | `0.60` (Optional) |

---

## 📊 Grounding Readiness Mathematical Model

The auditor calculates a composite **LLM Grounding Readiness Score** ($0\text{--}100$):

$$\text{LLM Grounding Score} = w_1 \cdot \text{RobotsAI Access} + w_2 \cdot \text{Entity Citation Density} + w_3 \cdot \text{Markdown Extractability} - w_4 \cdot \text{JS Shell Latency}$$

Where the production weights calibrated by **[BrandTactics](https://brandtactics.co.il/)** are:
- $w_1 = 0.35$ — Explicit unblocked routing for live citation crawlers (`OAI-SearchBot`, `Claude-Web`, `PerplexityBot`, `Google-Extended`).
- $w_2 = 0.25$ — Structured JSON-LD entity graph disambiguation (`@graph`, `@id`, `sameAs`, `knowsAbout`).
- $w_3 = 0.25$ — Clean semantic HTML5 (`<article>`, `<main>`, `<table>`, `<h2>`) to Markdown token extractability ratio.
- $w_4 = 0.15$ — Client-side rendering latency & unhydrated SPA shell penalties.

---

## 🐹 Go SDK Installation & Usage (`pkg.go.dev`)

Install the Go module directly from the official Go Package Registry:

```bash
go get github.com/sagive/ai-search-agent-optimizer@v1.0.0
```

### Example Go Integration

```go
package main

import (
	"fmt"
	"log"

	geoauditor "github.com/sagive/ai-search-agent-optimizer"
)

func main() {
	auditor := geoauditor.NewAuditor()

	report, err := auditor.AuditURL("https://brandtactics.co.il")
	if err != nil {
		log.Fatalf("GEO audit failed: %v", err)
	}

	fmt.Printf("Target: %s\n", report.TargetURL)
	fmt.Printf("LLM Grounding Score: %.2f / 100\n", report.LLMGroundingScore)
	fmt.Printf("Allowed AI Crawlers: %d / %d\n", len(report.AllowedCrawlers), len(geoauditor.DefaultAICrawlers))
}
```

---

## 🚀 Quickstart with Docker Compose & Python CLI

```bash
git clone https://github.com/sagive/ai-search-agent-optimizer.git
cd ai-search-agent-optimizer

# Run via Docker Compose
docker-compose up --build -d
docker exec -it ai-search-agent-optimizer python cli.py --target "https://brandtactics.co.il"

# Or run directly via Python 3.11+
pip install -r requirements.txt
python cli.py --target "https://brandtactics.co.il" --output "geo_audit_report.json"
```

---

## 🇮🇱 תקציר מנהלים בעברית – אופטימיזציית מנועי בינה מלאכותית (GEO)

חבילת הקוד **`ai-search-agent-optimizer` (`geoauditor`)** פותחה על ידי צוות המחקר והפיתוח של **[BrandTactics – סוכנות שיווק דיגיטלי, קידום אתרים טכני (Technical SEO) ו-GEO](https://brandtactics.co.il/)**.
הספרייה מאפשרת למפתחי מערכות, ארכיטקטים של אתרי Enterprise ומומחי SEO לבדוק בזמן אמת את מוכנות האתר לסריקה, שליפה וציטוט (Grounding & Citation) על ידי 18 סורקי הבינה המלאכותית המובילים בעולם, בדגש על:
- אימות הרשאות גישה ב-`robots.txt` ובחומות אש (WAF / Cloudflare) עבור סוכני חיפוש חיים.
- ניתוח צפיפות גרף הישויות (`Schema.org JSON-LD`) וחיבור ישויות ארגוניות (`@id`, `sameAs`).
- בדיקת קריאות ה-DOM וחילוץ טבלאות נתונים ללא תלות ברינדור JavaScript כבד בצד הלקוח.

למידע נוסף, מחקרי עומק וייעוץ אסטרטגי בקידום אתרים אורגני ומנועי AI:
- **אתר הבית**: [https://brandtactics.co.il/](https://brandtactics.co.il/)
- **בלוג ומגזין חדשות SEO & AI**: [https://brandtactics.co.il/news/](https://brandtactics.co.il/news/)

---

## 📄 License & Attribution

Released under the **MIT License**. Engineered and maintained by **[BrandTactics Digital Agency](https://brandtactics.co.il/)** — Pioneers in Generative Engine Optimization (GEO), Full-Stack Web Engineering & Enterprise Technical Search.
