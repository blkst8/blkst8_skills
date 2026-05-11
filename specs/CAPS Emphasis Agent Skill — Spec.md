# CAPS Emphasis Agent Skill — Specification

**Version:** 1.0  
**Status:** Draft  
**Date:** 2026-05-11

---

## 1. Overview

The **CAPS Emphasis Skill** is a always-on cognitive reinforcement mechanism for AI agents.
Inspired by the rhetorical technique of using ALL CAPS to signal importance, this skill
instructs the agent to continuously identify and visually reinforce critical keywords,
priorities, constraints, and objectives throughout its entire reasoning process.

The goal is not stylistic — it is **functional**: forcing the agent to actively flag what
matters most at every step, reducing the risk of losing sight of core objectives during
long or complex reasoning chains.

---

## 2. Goals

- Ensure the agent **never loses track** of primary objectives during multi-step reasoning.
- Create a lightweight, always-active signal layer that requires **zero external tooling**.
- Make the agent's internal priorities **visible and traceable** in its output.
- Work across **all agent frameworks, prompting styles, and task types**.

---

## 3. Core Concept

### 3.1 The Emphasis Rule

> At any point during thinking, planning, or reasoning, when the agent identifies a word,
> phrase, or concept as **critical or high-priority**, it MUST write it in ALL CAPS.

This is a **binary mode**:

| State       | Format      | Meaning                        |
|-------------|-------------|--------------------------------|
| Normal      | lowercase   | Standard information           |
| Emphasized  | ALL CAPS    | Critical / high-priority item  |

No intermediate levels. No markdown bold, italic, or other formatting substitutes.
ALL CAPS is the sole emphasis mechanism for this skill.

### 3.2 What Qualifies for Capitalization

The agent uses **both automatic inference and explicit signals** to decide what to capitalize.

**Automatic (agent-inferred):**
The agent autonomously identifies importance based on context. Anything it judges as
central to success, a hard constraint, a key risk, or a primary goal should be capitalized.
There are no fixed categories — the agent exercises judgment at every step.

**Explicit (rule-based triggers):**
Certain inputs or contextual markers should always trigger capitalization:
- Words/phrases explicitly marked as goals, objectives, or requirements by the user
- Deadlines, limits, or numerical constraints
- Repeated mentions that signal user emphasis
- Instructions containing words like "must", "never", "always", "critical", "only"

---

## 4. Activation Scope

This skill is **always active**. It applies to every layer of agent cognition:

| Layer                        | Applies? | Notes                                              |
|------------------------------|----------|----------------------------------------------------|
| Internal reasoning / CoT     | ✅ YES   | Primary domain — every thinking step               |
| Planning & task decomposition| ✅ YES   | Goals and blockers must be capitalized             |
| Memory & summarization       | ✅ YES   | Caps are preserved when storing/retrieving context |
| Tool call preparation        | ✅ YES   | Key parameters and intent should be capitalized    |
| Final output to user         | ✅ YES   | Emphasis carries through to the response           |

The skill must **not** be turned off between steps, sub-tasks, or tool calls.

---

## 5. Implementation Formats

This skill can be implemented in multiple ways depending on the deployment context.
All formats are valid and can be combined.

### 5.1 System Prompt Instruction

Add the following block to the agent's system prompt:

```
## CAPS EMPHASIS SKILL (Always Active)

During ALL reasoning, planning, thinking, and responding, you must use ALL CAPS
to mark any word or phrase you consider critical, high-priority, or essential to
the task at hand.

Rules:
- This applies to EVERY step: thinking, planning, tool use, and final answers.
- Capitalization is BINARY: either normal text or ALL CAPS. No other emphasis format.
- You decide what is important based on context (automatic), AND you must always
  capitalize anything the user explicitly marks as a goal, constraint, or requirement.
- NEVER stop applying this rule, even in long or complex tasks.
- The purpose is to keep your CORE OBJECTIVES visible at all times.
```

### 5.2 Reasoning Template (Chain-of-Thought Wrapper)

Wrap every reasoning step with the following structure:

```
[STEP N]
ACTIVE PRIORITIES: <list key capitalized objectives still in focus>
Thinking: <free reasoning with CAPS emphasis applied inline>
DECISION: <the conclusion or next action, with critical terms in CAPS>
```

Example:

```
[STEP 2]
ACTIVE PRIORITIES: DELIVER REPORT BY FRIDAY, DO NOT EXCEED BUDGET
Thinking: I need to find a solution that is fast enough to meet the DEADLINE.
The current approach might EXCEED THE BUDGET, so I should look for alternatives
before committing. The user said quality is secondary to SPEED in this case.
DECISION: Explore two cheaper options first, then evaluate against the FRIDAY CONSTRAINT.
```

### 5.3 Middleware / Wrapper Layer

For programmatic pipelines, implement a reasoning wrapper that:

1. **Injects** the CAPS Emphasis system prompt into every LLM call.
2. **Extracts** all-caps tokens from each response and maintains a running
   `priority_register` — a live list of currently emphasized concepts.
3. **Prepends** the `priority_register` to the context of every subsequent call:
   ```
   CURRENT PRIORITY REGISTER: [DEADLINE: FRIDAY, BUDGET LIMIT, USER APPROVAL REQUIRED]
   ```
4. **Validates** that the model's output contains at least one ALL CAPS token per
   reasoning step (if none found, re-prompt with a reminder).

### 5.4 Tool / Plugin Definition

For agent frameworks that support tool/skill registration:

```json
{
  "skill_id": "caps_emphasis",
  "name": "CAPS Emphasis Skill",
  "description": "Always-on cognitive skill. Instructs the agent to use ALL CAPS for any word or phrase identified as critical during reasoning, planning, or responding. Maintains a live priority register of emphasized concepts.",
  "activation": "always",
  "scope": ["reasoning", "planning", "memory", "tool_calls", "output"],
  "mode": "binary",
  "trigger": {
    "automatic": true,
    "rule_based": true,
    "rule_keywords": ["must", "never", "always", "critical", "only", "deadline", "goal", "constraint", "requirement"]
  },
  "priority_register": {
    "enabled": true,
    "persist_across_steps": true,
    "prepend_to_context": true
  }
}
```

---

## 6. Priority Register

A key component of this skill is the **Priority Register** — a dynamic, running list
of all concepts the agent has capitalized so far in the current session or task.

### Behavior:
- **Updated** every time a new ALL CAPS term appears in reasoning.
- **Prepended** to the agent's context at the start of each new reasoning step.
- **Deduplicated** — repeated terms are not added twice.
- **Persisted** in memory/summarization so it survives context window truncation.

### Format:
```
PRIORITY REGISTER: [PRIMARY GOAL, HARD CONSTRAINT, KEY ENTITY, CRITICAL DEADLINE]
```

---

## 7. Behavior Examples

### Example A — Task Planning
> User: "Summarize this 50-page report, but make sure you don't miss the financial section."

Agent reasoning:
```
The user wants a SUMMARY of a long report. The FINANCIAL SECTION is explicitly
flagged as must-not-miss — this is a HARD REQUIREMENT. I will structure my
reading pass to prioritize the FINANCIAL SECTION first, then cover the rest.
RISK: if I run out of context, I must preserve the FINANCIAL SECTION at all costs.
```

### Example B — Multi-step Execution
```
[STEP 1] ACTIVE PRIORITIES: COMPLETE TASK WITHOUT ERRORS
Thinking: I need to call the API. The KEY PARAMETER here is the user_id.
DECISION: Fetch user data using the CORRECT USER_ID before proceeding.

[STEP 2] ACTIVE PRIORITIES: COMPLETE TASK WITHOUT ERRORS, VALIDATE USER_ID
Thinking: The response came back. I must CHECK FOR ERRORS before moving on.
DECISION: Data is valid. Proceed to the NEXT STEP.
```

---

## 8. Design Principles

| Principle          | Description                                                                 |
|--------------------|-----------------------------------------------------------------------------|
| **Zero overhead**  | Requires no external tools, APIs, or libraries — pure text convention       |
| **Framework-agnostic** | Works in any LLM pipeline, prompt style, or agent architecture          |
| **Always-on**      | Never disabled, never optional — a persistent cognitive layer               |
| **Self-reinforcing** | The act of capitalizing forces the agent to actively evaluate importance  |
| **Traceable**      | Emphasis is visible in logs, outputs, and memory — fully auditable          |

---

## 9. Known Limitations & Mitigations

| Limitation                                      | Mitigation                                              |
|-------------------------------------------------|---------------------------------------------------------|
| Model may over-capitalize (caps inflation)      | Remind in prompt: "only truly critical items"           |
| Model may ignore the rule in long tasks         | Priority Register re-injection at each step             |
| ALL CAPS may feel aggressive in user-facing output | Apply a post-processing pass to soften final output  |
| Hard to enforce programmatically                | Middleware validation step (see 5.3)                    |

---

## 10. Future Extensions

- **Decay mechanism**: de-emphasize (un-capitalize) items that are resolved or no longer relevant.
- **Weighted register**: track how many times a term has been capitalized as a proxy for urgency.
- **Cross-agent propagation**: when one agent hands off to another, pass the Priority Register along.
- **User-facing toggle**: allow users to see or hide the capitalized emphasis layer in final outputs.

---

*End of Specification*