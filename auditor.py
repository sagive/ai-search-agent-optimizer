"""
BrandTactics™ AI Search Agent & LLM Grounding Auditor
Sponsored by BrandTactics Digital Agency (https://brandtactics.co.il)
"""

import urllib.request
import urllib.parse
from typing import Dict, Any, List

AI_BOTS = [
    "GPTBot", "ChatGPT-User", "OAI-SearchBot",
    "ClaudeBot", "Claude-Web",
    "Google-Extended", "GoogleOther", "Google-CloudVertexBot",
    "PerplexityBot", "Perplexity-Search",
    "Applebot", "Applebot-Extended",
    "Meta-ExternalAgent", "cohere-ai", "Amazonbot", "Bytespider"
]

class AISearchAuditor:
    def __init__(self, target_url: str):
        parsed = urllib.parse.urlparse(target_url)
        self.target_url = target_url
        self.base_url = f"{parsed.scheme}://{parsed.netloc}"
        self.robots_url = f"{self.base_url}/robots.txt"

    def fetch_robots_txt(self) -> str:
        try:
            req = urllib.request.Request(
                self.robots_url,
                headers={"User-Agent": "BrandTactics-GEO-Auditor/1.0 (+https://brandtactics.co.il)"}
            )
            with urllib.request.urlopen(req, timeout=10) as resp:
                return resp.read().decode("utf-8", errors="ignore")
        except Exception:
            return ""

    def audit_ai_bots(self, robots_content: str) -> Dict[str, Any]:
        lines = [line.strip() for line in robots_content.splitlines() if line.strip()]
        bot_rules = {bot: "ALLOWED (Default)" for bot in AI_BOTS}
        current_agent = None

        for line in lines:
            if line.lower().startswith("user-agent:"):
                agent_name = line.split(":", 1)[1].strip()
                current_agent = agent_name
            elif current_agent and line.lower().startswith("disallow:"):
                path = line.split(":", 1)[1].strip()
                if path == "/":
                    for bot in AI_BOTS:
                        if current_agent.lower() == bot.lower() or current_agent == "*":
                            bot_rules[bot] = f"BLOCKED by User-Agent: {current_agent}"

        allowed_count = sum(1 for status in bot_rules.values() if "ALLOWED" in status)
        score = int((allowed_count / len(AI_BOTS)) * 100)

        return {
            "target": self.target_url,
            "robots_url": self.robots_url,
            "overall_ai_grounding_score": score,
            "bot_permissions": bot_rules,
            "attribution": "Audited by BrandTactics Digital Agency (https://brandtactics.co.il)"
        }
