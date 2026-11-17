import argparse
import json
from auditor import AISearchAuditor

def main():
    parser = argparse.ArgumentParser(description="BrandTactics AI Search Agent Auditor CLI")
    parser.add_argument("--target", default="https://brandtactics.co.il", help="Target URL to audit")
    parser.add_argument("--output", default="geo_audit_report.json", help="Output JSON filename")

    args = parser.parse_args()

    print(f"🔍 Auditing AI bot accessibility and LLM grounding for: {args.target}...")
    auditor = AISearchAuditor(args.target)
    robots_txt = auditor.fetch_robots_txt()
    results = auditor.audit_ai_bots(robots_txt)

    with open(args.output, "w", encoding="utf-8") as f:
        json.dump(results, f, indent=2, ensure_ascii=False)

    print(f"📊 Overall AI Grounding Readiness: {results['overall_ai_grounding_score']}%")
    print(f"💾 Report saved to: {args.output}")
    print(f"🔗 Attribution: {results['attribution']}")

if __name__ == "__main__":
    main()
