# auto idea — Brainstorm and Evaluate Ideas

> [!NOTE]
> Global Connectivity: [Registry](../SKILL.md) | [Setup](setup.md) | [Plan](plan.md)

Brainstorm ideas using multi-provider orchestra, evaluate with ICE scoring, and save as BS file.

## Flags

| Flag | Description |
|------|-------------|
| `--strategy` | Orchestra strategy: debate (default), consensus, pipeline, fastest |
| `--providers` | Provider list override (default: all configured) |

## Pipeline

### [REQUIRED] Step 1: Parse Input

Extract from remaining arguments:
- `IDEA_DESC`: description
- `STRATEGY`: default debate
- `PROVIDERS`: default all

### [REQUIRED] Step 2: Structure as What/Why/Who/When

Structure the idea description into What, Why, Who, When.

### [REQUIRED] Step 3: Call Orchestra Brainstorm

```bash
auto orchestra brainstorm "{structured idea}" --strategy {STRATEGY}
```

### [REQUIRED] Step 4: ICE Scoring

Score Impact, Confidence, Ease (1-10).
`Score = (Impact × Confidence × Ease) / 100`

### [REQUIRED] Step 5: Save BS File

Save to `.autopus/brainstorms/BS-{ID}.md`.
Display workflow lifecycle bar: `● idea → ○ plan → ○ go → ○ sync`
