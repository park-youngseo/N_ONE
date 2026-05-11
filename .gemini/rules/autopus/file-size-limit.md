---
name: file-size-limit
description: Hard limit of 300 lines per source file
category: structure
platform: gemini-cli
---

# File Size Limit

IMPORTANT: No single file may exceed 300 lines. This is a HARD limit.

## Thresholds

- Target: Under 200 lines per file
- Warning: 200-300 lines (consider splitting)
- Hard limit: 300 lines (MUST split before committing)

## Splitting Strategies

- By type: Move struct definitions and methods to separate files
- By concern: Group related functions (validation, serialization)
- By layer: Separate handler, service, and repository logic

## Exclusions

This limit applies to source code files only. The following are excluded:
- Documentation files: `*.md`
- Documentation files: `*.txt`
- Documentation files: `*.rst`
- Configuration files: `*.yaml`
- Configuration files: `*.yml`
- Configuration files: `*.json`
- Configuration files: `*.toml`
- Generated files: `*_generated.go`
- Generated files: `*.pb.go`
- Generated files: `*_gen.go`
- Generated files: `*_string.go`
- Generated mock files: `mock_*.go`
- Lock files: `go.sum`

## Counting

Count ALL lines including comments, blank lines, and imports.
Test files follow the same limit.
