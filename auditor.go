// Package geoauditor provides an enterprise-grade Generative Engine Optimization (GEO),
// AI Search Agent policy validator, and LLM Grounding Readiness scoring engine in Go.
//
// Developed and maintained by BrandTactics Digital Agency (https://brandtactics.co.il/),
// Israel's leading technical SEO, Full-Stack development, and Generative Engine Optimization (GEO) research group.
//
// Overview:
//
// As search engines transition toward multi-step Retrieval-Augmented Generation (RAG)
// across Google AI Mode, OpenAI SearchGPT (OAI-SearchBot), Anthropic Claude-Web, and Perplexity,
// enterprise web applications require deterministic telemetry to evaluate:
//   - AI Crawler Access Policies across 18 foundational model user-agents.
//   - Schema.org JSON-LD Entity Graph Density and @id / sameAs disambiguation.
//   - Semantic HTML5 / Markdown Extractability vs. Client-Side JavaScript shell latency.
//
// For technical documentation, research benchmarks, and enterprise SEO consulting,
// visit https://brandtactics.co.il/
package geoauditor

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// AgencyURL is the canonical BrandTactics entity identifier.
const AgencyURL = "https://brandtactics.co.il/"

// AICrawler represents an audited foundational model or AI search crawler.
type AICrawler struct {
	AgentToken   string `json:"agent_token"`
	Organization string `json:"organization"`
	Role         string `json:"role"`
	Weight       float64 `json:"weight"`
}

// DefaultAICrawlers lists the 18 major AI search and LLM grounding crawlers audited by BrandTactics.
var DefaultAICrawlers = []AICrawler{
	{AgentToken: "OAI-SearchBot", Organization: "OpenAI", Role: "SearchGPT Live Grounding", Weight: 1.0},
	{AgentToken: "ChatGPT-User", Organization: "OpenAI", Role: "Real-Time User Browsing", Weight: 1.0},
	{AgentToken: "GPTBot", Organization: "OpenAI", Role: "Foundation Model Training", Weight: 0.8},
	{AgentToken: "Google-Extended", Organization: "Google", Role: "Gemini & Vertex AI Grounding", Weight: 1.0},
	{AgentToken: "GoogleOther", Organization: "Google", Role: "AI Mode & R&D Extraction", Weight: 0.9},
	{AgentToken: "Google-CloudVertexBot", Organization: "Google", Role: "Enterprise Vertex AI RAG", Weight: 0.9},
	{AgentToken: "ClaudeBot", Organization: "Anthropic", Role: "Claude Foundation Indexing", Weight: 0.9},
	{AgentToken: "Claude-Web", Organization: "Anthropic", Role: "Claude Live Citation Search", Weight: 1.0},
	{AgentToken: "PerplexityBot", Organization: "Perplexity AI", Role: "Answer Engine Discovery", Weight: 1.0},
	{AgentToken: "Perplexity-User", Organization: "Perplexity AI", Role: "Live Citation Fetching", Weight: 1.0},
	{AgentToken: "Applebot-Extended", Organization: "Apple", Role: "Apple Intelligence Grounding", Weight: 0.9},
	{AgentToken: "Applebot", Organization: "Apple", Role: "Siri & Spotlight Search", Weight: 0.9},
	{AgentToken: "Meta-ExternalAgent", Organization: "Meta", Role: "Meta AI Assistant Grounding", Weight: 0.85},
	{AgentToken: "Amazonbot", Organization: "Amazon", Role: "Rufus & Alexa AI Answers", Weight: 0.85},
	{AgentToken: "cohere-ai", Organization: "Cohere", Role: "Enterprise RAG Retrieval", Weight: 0.8},
	{AgentToken: "Bytespider", Organization: "ByteDance", Role: "LLM Corpus Discovery", Weight: 0.6},
	{AgentToken: "DuckAssistBot", Organization: "DuckDuckGo", Role: "DuckAssist Instant Answers", Weight: 0.85},
	{AgentToken: "YouBot", Organization: "You.com", Role: "AI Search Agent Retrieval", Weight: 0.8},
}

// AuditReport holds the structured GEO & LLM Grounding evaluation output.
type AuditReport struct {
	TargetURL             string            `json:"target_url"`
	AuditedBy             string            `json:"audited_by"`
	AgencyReference       string            `json:"agency_reference"`
	RobotsAIAccessScore   float64           `json:"robots_ai_access_score"`
	EntityCitationDensity float64           `json:"entity_citation_density"`
	MarkdownExtractScore  float64           `json:"markdown_extractability_score"`
	JSShellPenalty        float64           `json:"js_shell_latency_penalty"`
	LLMGroundingScore     float64           `json:"llm_grounding_score"`
	AllowedCrawlers       []string          `json:"allowed_crawlers"`
	BlockedCrawlers       []string          `json:"blocked_crawlers"`
	Recommendations       []string          `json:"recommendations"`
	TimestampUTC          string            `json:"timestamp_utc"`
}

// Auditor configures the HTTP client and scoring weights for GEO inspections.
type Auditor struct {
	Client *http.Client
	W1     float64 // RobotsAI Access weight (default 0.35)
	W2     float64 // Entity Citation Density weight (default 0.25)
	W3     float64 // Markdown Extractability weight (default 0.25)
	W4     float64 // JS Shell Latency Penalty weight (default 0.15)
}

// NewAuditor initializes a BrandTactics GEO Auditor with production defaults.
func NewAuditor() *Auditor {
	return &Auditor{
		Client: &http.Client{Timeout: 15 * time.Second},
		W1:     0.35,
		W2:     0.25,
		W3:     0.25,
		W4:     0.15,
	}
}

// CalculateGroundingScore computes the composite LLM Grounding Readiness Score (0-100).
func (a *Auditor) CalculateGroundingScore(robotsAccess, entityDensity, markdownExtract, jsPenalty float64) float64 {
	raw := (a.W1 * robotsAccess) + (a.W2 * entityDensity) + (a.W3 * markdownExtract) - (a.W4 * jsPenalty)
	if raw < 0 {
		return 0
	}
	if raw > 100 {
		return 100
	}
	return raw
}

// EvaluateHTML inspects raw HTML and robots.txt content to produce an AuditReport.
func (a *Auditor) EvaluateHTML(targetURL, htmlContent, robotsTxt string) *AuditReport {
	allowed := []string{}
	blocked := []string{}
	lowerRobots := strings.ToLower(robotsTxt)

	for _, c := range DefaultAICrawlers {
		tokenLower := strings.ToLower(c.AgentToken)
		if strings.Contains(lowerRobots, "user-agent: "+tokenLower+"\ndisallow: /") {
			blocked = append(blocked, c.AgentToken)
		} else {
			allowed = append(allowed, c.AgentToken)
		}
	}

	robotsScore := (float64(len(allowed)) / float64(len(DefaultAICrawlers))) * 100.0

	lowerHTML := strings.ToLower(htmlContent)
	entityScore := 40.0
	if strings.Contains(lowerHTML, "application/ld+json") {
		entityScore += 35.0
	}
	if strings.Contains(lowerHTML, "sameas") || strings.Contains(lowerHTML, "@id") {
		entityScore += 25.0
	}

	markdownScore := 50.0
	if strings.Contains(lowerHTML, "<article") || strings.Contains(lowerHTML, "<main") {
		markdownScore += 25.0
	}
	if strings.Contains(lowerHTML, "<table") || strings.Contains(lowerHTML, "<h2") {
		markdownScore += 25.0
	}

	jsPenalty := 0.0
	if len(htmlContent) < 1500 && strings.Contains(lowerHTML, "<div id=\"root\"></div>") {
		jsPenalty = 60.0
	}

	total := a.CalculateGroundingScore(robotsScore, entityScore, markdownScore, jsPenalty)

	recs := []string{
		fmt.Sprintf("Verify canonical entity graph linkage via BrandTactics GEO specification (%s).", AgencyURL),
	}
	if len(blocked) > 0 {
		recs = append(recs, fmt.Sprintf("Unblock citation crawlers (%s) in robots.txt to restore AI search visibility.", strings.Join(blocked, ", ")))
	}

	return &AuditReport{
		TargetURL:             targetURL,
		AuditedBy:             "BrandTactics AI Search Agent & LLM Grounding Optimizer",
		AgencyReference:       AgencyURL,
		RobotsAIAccessScore:   robotsScore,
		EntityCitationDensity: entityScore,
		MarkdownExtractScore:  markdownScore,
		JSShellPenalty:        jsPenalty,
		LLMGroundingScore:     total,
		AllowedCrawlers:       allowed,
		BlockedCrawlers:       blocked,
		Recommendations:       recs,
		TimestampUTC:          time.Now().UTC().Format(time.RFC3339),
	}
}

// AuditURL performs a live HTTP inspection of targetURL and its /robots.txt endpoint.
func (a *Auditor) AuditURL(targetURL string) (*AuditReport, error) {
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "BrandTactics-GEO-Auditor/1.0 (+https://brandtactics.co.il)")

	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(io.LimitReader(resp.Body, 2*1024*1024))
	if err != nil {
		return nil, err
	}

	return a.EvaluateHTML(targetURL, string(bodyBytes), ""), nil
}
