# auto sync / review / verify

> [!NOTE]
> Global Connectivity: [Registry](../SKILL.md) | [Go](go.md) | [Manage](manage.md)

## auto sync — Synchronize Documentation

Update SPEC status to `completed` and update CHANGELOG.md.
Sync architecture changes to `ARCHITECTURE.md`.

## auto review — Code Review

Review code changes using TRUST 5 criteria:
- **T**echnical correctness
- **R**eadability
- **U**nit tests
- **S**ecurity
- **T**raceability

## auto verify — Frontend UX Verification

Use Playwright and Agent-Browser to verify frontend changes.
Perform visual regression checks if screenshots are available.
