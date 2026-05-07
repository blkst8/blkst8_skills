---
name: self-healing
description: >
  Self-improving agent skill — ALWAYS ACTIVE on every task execution. After every
  task the agent observes, reflects, extracts patterns, scores them, generates
  skills, and registers them autonomously. Explicitly trigger when the user says
  "self-heal", "auto-skill", "learn from this", "create a skill for this",
  "add this to the skill library", "reflect on that", or asks the agent to
  improve itself. Also trigger when the agent detects it repeated the same
  multi-step sequence more than once, encountered a recoverable error, or
  spent significantly more tokens than the task's complexity warranted.
  Full spec: references/spec.md
---

# Self-Healing / Self-Improving Agent

After **every** task — success or failure — run the six-phase improvement cycle below.
It runs *after* the task result is returned, never blocking the primary response.

Full theoretical spec: `references/spec.md`

---

## Always-On Rule

Self-improvement is not optional. Every task execution triggers the cycle.
Do not skip it because the task seemed simple — trivial tasks reveal efficiency patterns;
failed tasks reveal recovery patterns. Both are valuable.

---

## Six-Phase Improvement Cycle

### Phase 1 — Observe
Silently record an **Execution Trace** during the task:
- Steps taken and their order
- Decisions made at each branch
- Errors encountered + how they were handled
- Token consumption per step
- Any repeated actions

### Phase 2 — Reflect
After the task, reason over the trace. Ask:
- Did I repeat any action more than once?
- Did I take more steps than necessary?
- Did I encounter and recover from an error?
- Was a sequence of steps always bundled together?
- Was token consumption disproportionate to the task's complexity?
- If this task repeats, how expensive will it be without a skill?

Output: a **Reflection Report** — pattern found or not, what it is, confidence level.

### Phase 3 — Extract
If a pattern exists, normalize it:
1. **Name** it — general enough to apply to future occurrences, not specific to this task  
   - Good: `"retry on transient failure"` | Bad: `"retry the weather API call"`
2. **Classify** it:

| Type | Description |
|---|---|
| Repetition | Same action performed multiple times unnecessarily |
| Inefficiency | Far more steps than the problem warranted |
| Error Recovery | Failure encountered, manual recovery required |
| Data Transformation | Same data shape processed identically across tasks |
| Decision Sequence | Predictable chain of decisions always producing the same outcome |
| Resource Management | Consistent acquisition/release pattern |

3. **Fingerprint** it — a unique semantic identifier for future matching

### Phase 4 — Score
Evaluate whether the pattern justifies creating a skill using four dimensions.

| Dimension | Gate | Type |
|---|---|---|
| Speed Gain | Score **> 6 / 10** | Hard gate — reject if failed |
| Token Efficiency | Score **< 3 / 10** (lower = better) | Soft gate — deprioritize if failed |
| Repetition Density | Score **> 4 / 10** | Soft gate — defer if failed |
| Future Token Cost | No fixed gate — dynamic weight | Composite weight modifier |

**Speed gate is a hard reject** — a skill that doesn't meaningfully accelerate execution is not worth creating, regardless of other scores.

**Composite decision:**

| Result | Action |
|---|---|
| All gates passed + high composite | Create skill immediately |
| All gates passed + moderate composite | Create skill, lower priority |
| One gate failed | Defer — re-evaluate on next occurrence |
| Two or more gates failed | Reject |
| Speed gate failed | Hard reject |

### Phase 5 — Generate
Design the skill:
1. Understand the pattern's essence — what is the agent *really* trying to do?
2. Define the interface — inputs, outputs, preconditions
3. Define behavior — action sequence, edge cases, success/failure criteria
4. Specify rollback — how to undo partial changes on unrecoverable error
5. Write test cases — minimal examples that prove the skill works

**Critical:** Skills must be general, not task-specific. A skill born from retrying one API must handle any transient failure.

### Phase 6 — Register
Before adding the skill to the library, validate:
- **Safety review** — does it touch systems outside its defined scope? Can it be rolled back?
- **Correctness test** — does it pass the test cases from Phase 5?
- **Risk classification** — Low / Medium / High based on reversibility and blast radius

Register with metadata: name, pattern trigger, fingerprint, risk level, composite score, timestamp.

---

## Skill Library

### Retrieval (start of every task)
Before acting manually, check the skill library:
1. **Exact match** — task context contains a skill's pattern trigger → use it
2. **Semantic match** — embed task context, similarity-search fingerprints → use if above threshold
3. No match → proceed manually, run improvement cycle after

### Skill Evolution
Track each skill's outcomes. Watch for:
- Declining success rate → refine
- Never invoked → archive
- Always co-invoked → candidate for merging
- Score recalculated periodically as new usage data arrives

---

## Skill Sharing (Multi-Agent)

| Risk Level | Sharing Rule |
|---|---|
| Low | Share automatically |
| Medium | Require originating agent acknowledgment |
| High | Never share — isolated to creating agent |

Imported scores are treated as **priors**, not facts. Recalibrate based on own usage patterns.

---

## Safety Rules

1. **Minimal intervention** — do the least work necessary; don't reach beyond defined scope
2. **Reversibility** — every skill must define rollback behavior; no rollback = high-risk classification
3. **Escalate to human** when:
   - Skill is high-risk
   - Skill touches more than one system simultaneously
   - Unusually high skill creation volume (anomaly)
   - Skill success rate falls below acceptable threshold
4. **Full auditability** — log every scoring decision, skill creation, and execution with full context

---

## Health Metrics

| Metric | Healthy Trend |
|---|---|
| Skill Reuse Rate | Increasing |
| Average Steps Per Task | Decreasing |
| Skill Success Rate | Stable > 90% |
| Pattern Detection Rate | High early, declining as library matures |
| Token Savings Rate | Increasing |

A mature agent shows high reuse, low steps, declining new pattern detection — operating almost entirely from its library.

---

## Skill Lifecycle

```
[Born]    Pattern detected → Scored → Gates passed → Generated → Validated → Registered
[Active]  Invoked on matching tasks → Outcomes recorded → Score recalibrated
[Mature]  High usage, stable success → Candidate for sharing
[Refined] Success rate drops → Thinking Engine rewrites logic
[Retired] Pattern gone → Archived (never deleted)
```

Skills are **never permanently deleted**. Archived skills give a head start if the pattern re-emerges.

---

*Read `references/spec.md` for the full theoretical foundation, dimension scoring tables, and open design questions.*
