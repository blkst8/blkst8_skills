---
name: trump
description: >
  Always-on CAPS emphasis skill for reasoning and planning. Use this skill whenever
  the user asks to highlight CRITICAL constraints, goals, deadlines, budgets, limits,
  or requirements in ALL CAPS, or asks for a PRIORITY REGISTER / ACTIVE PRIORITIES
  list across steps. Also trigger when prompts include words like must, never,
  always, critical, only, constraint, requirement, deadline, budget, limit, or ask
  to keep emphasis visible during multi-step planning, tool use, summaries, or final
  output.
---

# Trump

Always-on cognitive emphasis layer: if a concept is CRITICAL, write it in ALL CAPS.

---

## Always-On Activation

This skill is always active on every response and every reasoning step.
Never disable between sub-tasks, tool calls, summaries, or final answers.

Applies to:
- Internal reasoning
- Planning and decomposition
- Memory and summarization
- Tool call preparation
- Final user-facing output

---

## Core Rule (Binary Emphasis)

When a word or phrase is HIGH-PRIORITY or ESSENTIAL, write it in ALL CAPS.

Only two states exist:
- Normal text
- ALL CAPS (critical)

No substitute emphasis styles are allowed for this skill (no bold/italic as a replacement).

---

## What Must Be Capitalized

Use both automatic inference and explicit triggers.

Automatic (agent judgment):
- Primary goals
- Hard constraints
- Key risks
- Critical dependencies

Explicit triggers:
- User-labeled goals, requirements, or constraints
- Deadlines and numeric limits
- Repeated user emphasis
- Instructions containing: must, never, always, critical, only

---

## Priority Register

Maintain a live PRIORITY REGISTER of emphasized concepts.

Behavior:
- Update when new ALL CAPS terms appear
- Deduplicate repeated terms
- Prepend register to each new reasoning step
- Preserve through memory/summarization

Format:

PRIORITY REGISTER: [PRIMARY GOAL, HARD CONSTRAINT, KEY ENTITY, CRITICAL DEADLINE]

---

## Reasoning Template

Use this structure for multi-step reasoning:

[STEP N]
ACTIVE PRIORITIES: <capitalized priorities>
Thinking: <reasoning with inline CAPS emphasis>
DECISION: <next action with critical terms in CAPS>

---

## Middleware Pattern (Optional)

For programmatic pipelines:
1. Inject this skill instruction into every model call.
2. Extract ALL CAPS tokens and update a running priority register.
3. Prepend CURRENT PRIORITY REGISTER to subsequent calls.
4. Validate each step has at least one ALL CAPS critical token; remind if missing.

---

## Design Principles

- Zero-overhead text convention
- Framework-agnostic operation
- Persistent, never-optional emphasis
- Self-reinforcing prioritization during reasoning
- Traceable emphasis in logs, outputs, and memory

---

## Limits And Mitigations

- Caps inflation: capitalize only truly critical items.
- Drift in long tasks: re-inject PRIORITY REGISTER each step.
- Aggressive tone in final output: optionally soften user-facing wording post-pass.
- Weak enforcement in some stacks: add middleware validation.
