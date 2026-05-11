# auto go — Implement a SPEC

> [!NOTE]
> Global Connectivity: [Registry](../SKILL.md) | [Plan](plan.md) | [Sync](sync.md)

Implement code based on a SPEC. Follows TDD methodology.

## Pipeline

### Step 0: Parse Flags

Extract `SPEC-ID` and global flags.

### Step 1: Initialize Pipeline (MUST call Agent tool)

Spawn an `executor` agent to:
1. Locate `SPEC_PATH`.
2. Analyze codebase for target files.
3. Determine `WORKING_DIR`.

### Step 2: Implement Tests (TDD)

Write tests first in the `WORKING_DIR`.
Verify tests FAIL before implementation.

### Step 3: Implement Code

Implement logic to satisfy tests.
Ensure files do not exceed 300 lines (AX-300).

### Step 4: Verify & Fix

Run tests and fix issues until PASS.

### Step 5: Completion

Display workflow lifecycle bar: `● go → ○ sync`
Guide to `/auto sync`.
