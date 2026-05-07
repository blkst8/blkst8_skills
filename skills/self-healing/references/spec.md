# Self-Improving Agent — Spec
> v1.1.0 · 2026-05-07 · Draft

---

## 1. Overview

Agent that watches own execution, finds recurring patterns, turns 'em into reusable skills. Improves without retraining — pure runtime skill accumulation.

> *Never solve same problem twice same way.*

Not fine-tuning. Not retraining. **Runtime self-improvement** — agent grows smarter within operational lifetime, no human intervention.

**Always-On:** Self-improvement's not optional. Every task — success or failure — runs through reflection + pattern-detection pipeline. Always watching.

---

## 2. Three Pillars

**Pillar 1 — Metacognition:** After every task, ask: *"How'd I do that? Could've been faster? Done this before?"* Self-awareness seeds all improvement.

**Pillar 2 — Pattern Abstraction:** Abstract specific struggles into general rules. Not *"retry weather API"* — *"retry on transient failure."*

**Pillar 3 — Skill Encoding:** Pattern → skill. Skills aren't descriptions; they're capabilities. Invokable, measurable, improvable, shareable.

**Feedback loop:**
```
Experience → Reflection → Pattern → Skill Scoring → Skill → Better Experience → ...
```

Each cycle = measurably more capable agent. Library grows → less manual effort per task.

**Why it works:** Classical ML changes weights — slow, expensive, offline. This changes **available actions**. Each skill = richer vocabulary. Intelligence grows without becoming different model.

---

## 3. Core Concepts

| Concept | Def |
|---|---|
| **Task** | Unit of work agent's asked to do |
| **Execution Trace** | Full record of every step, decision, outcome |
| **Reflection** | Structured reasoning over execution trace |
| **Pattern** | Recurring, nameable behavior across traces |
| **Skill** | Named, reusable capability from pattern — stored + invocable |
| **Skill Score** | Composite rating: worth creating? how urgent? |
| **Skill Library** | Persistent, growing collection of all skills |
| **Thinking Engine** | LLM-powered reasoning module for reflection + skill design |
| **Pattern Fingerprint** | Unique semantic ID for pattern — matches future occurrences |

---

## 4. Six-Phase Cycle

Runs after every task execution.

---

### Phase 1 — Observe

Silent **Execution Trace** collected during task:
- Steps + order
- Decisions at each branch
- Inputs consumed, outputs produced
- Errors + recovery
- Time per step
- Repeated actions
- Tokens consumed

Agent doesn't modify behavior for tracing — acts normally, trace collected passively. Performance never degraded.

**Trace = raw material for all learning.**

---

### Phase 2 — Reflect

Thinking Engine examines trace, asks:
- *Repeated any action?*
- *More steps than needed?*
- *Hit error, had to recover?*
- *Sequence always appeared together?*
- *Token use proportionate to complexity?*
- *How expensive if task repeats without skill?*

Not rule-check — **reasoning process**. Reads trace like journal entry, finds themes + inefficiencies.

Output: **Reflection Report** — pattern found/not, what it is, confidence.

**Always triggered.** Simple task → efficiency pattern. Failed task → recovery pattern. Both valuable.

---

### Phase 3 — Extract

Pattern found → normalize into reusable definition. Three ops:

**1. Name** — general, not task-specific:
- Good: `"retry on transient failure"`
- Bad: `"retry the weather API call"`

**2. Classify:**

| Type | Desc |
|---|---|
| Repetition | Same action performed multiple times unnecessarily |
| Inefficiency | Far more steps than problem warranted |
| Error Recovery | Failure hit, manual recovery required |
| Data Transformation | Same data shape processed identically, repeatedly |
| Decision Sequence | Predictable chain → same outcome |
| Resource Management | Consistent acquire/release pattern |

**3. Fingerprint** — semantic vector of normalized description. Used to match future occurrences even when surface wording differs.

---

### Phase 4 — Score

> *Worth creating skill for this right now?*

Four scoring dimensions → **Composite Skill Score** → create or not.

---

#### Dim 1 — Speed Gain
*How much faster will future tasks be?*

Gate: **score > 6 / 10** — hard gate. Fails → reject regardless of other scores.

| Score | Meaning |
|---|---|
| 1–3 | Negligible — task already fast |
| 4–6 | Moderate — fails gate |
| 7–8 | Significant — noticeably faster |
| 9–10 | Dramatic — collapses many steps to one |

---

#### Dim 2 — Token Efficiency
*How much does skill cut manual token cost?*

Gate: **score < 3 / 10** — lower is better. Score ≥ 3 → deprioritize.

| Score | Meaning |
|---|---|
| 1–2 | Eliminates nearly all overhead — high savings |
| 3 | Borderline — fails gate |
| 4–6 | Manual approach already moderate — low benefit |
| 7–10 | Manual already lean — skill adds little |

---

#### Dim 3 — Repetition Density
*How often does pattern appear?*

Gate: **score > 4 / 10** — soft gate. Fails → defer.

| Score | Meaning |
|---|---|
| 1–2 | Once or twice — too rare |
| 3–4 | Occasionally — fails gate |
| 5–6 | Repeats regularly — meaningful use |
| 7–10 | Extremely frequent — urgently needed |

---

#### Dim 4 — Future Token Cost (Dynamic Weight)
*Total token cost of NOT creating skill across all future occurrences?*

No fixed gate — **dynamic weight**. High projected cost → this dim dominates composite. Low → near-irrelevant.

| Future Cost | Weight |
|---|---|
| Very High | Max — dominates score |
| High | Elevated — strong influence |
| Moderate | Standard |
| Low | Reduced |
| Very Low | Minimal |

---

#### Composite Decision

| Result | Action |
|---|---|
| All gates passed + high composite | **Create immediately** |
| All gates passed + moderate | **Create, lower priority** |
| One gate failed | **Defer — re-eval on next occurrence** |
| Two+ gates failed | **Reject** |
| Speed gate failed | **Hard reject** |

**Gate summary:**

| Dim | Gate | Type |
|---|---|---|
| Speed Gain | > 6 | Hard — reject if failed |
| Token Efficiency | < 3 | Soft — deprioritize if failed |
| Repetition Density | > 4 | Soft — defer if failed |
| Future Token Cost | none | Weight modifier only |

**Why asymmetric:** Speed = non-negotiable (hard gate). Token savings = strong signal, not absolute. Repetition = prerequisite for value. Future cost = most forward-looking — dynamic weight maximizes influence exactly when inaction's most costly.

---

### Phase 5 — Generate

Thinking Engine designs skill. Not template-filling — **creative reasoning**.

1. Understand pattern's essence — what's agent really trying to do?
2. Design interface — inputs, outputs, preconditions
3. Define behavior — action sequence, edge cases, success/failure criteria
4. Specify rollback — how to undo partial changes on error
5. Write test cases — minimal examples proving it works

Result: **Skill Blueprint** — ready for validation + registration.

**Critical:** Skills must be *general*. Skill born from one API retry must handle any transient failure — DB, file read, network. Never task-specific.

---

### Phase 6 — Register

Validation before library entry:

- **Safety review** — dangerous ops? Appropriate error handling? Blast radius limited?
- **Correctness test** — passes Phase 5 test cases?
- **Risk classification** — Low / Medium / High based on reversibility + scope

Registered with: name, pattern trigger, fingerprint, risk level, composite score, timestamp, perf stats.

From now on: matching pattern → skill invoked directly. No re-solving from scratch.

---

## 5. Skill Library

**Nature:** Agent's long-term procedural memory. Persists indefinitely — not context-window scoped. Never forgets a skill.

**Retrieval (start of every task):**
1. **Exact match** — task context matches pattern trigger → use it
2. **Semantic match** — embed context, similarity-search fingerprints → use if above threshold
3. No match → proceed manually → run improvement cycle after

**Evolution:** Track outcomes. Watch for:
- Declining success rate → refine
- Never invoked → obsolete candidate
- Always co-invoked → merge candidate
- Composite score recalculated periodically with new usage data

Second-order loop: not just creating skills, but improving existing ones.

---

## 6. Always-On

Improvement runs alongside every task — not a special mode.

**Why always-on:**
- **Consistency** — partial activation = unpredictable behavior. Operators can't reason about when agent's learning.
- **Completeness** — patterns don't announce themselves. Only way to catch all = observe all.
- **Compounding** — every-task reflector builds richer library. Difference vs. occasional reflector grows exponentially.

**Overhead:** Negligible. Reflection, extraction, scoring run *after* task. Skill retrieval = fast lookup. Primary execution never blocked.

---

## 7. Multi-Agent Sharing

Shared Skill Library → **collective intelligence**. Agent A creates skill → enters common pool → Agent B uses it without rediscovering.

**Trust rules:**

| Risk Level | Rule |
|---|---|
| Low | Share automatically |
| Medium | Needs originating agent acknowledgment |
| High | Never shared — isolated to creator |

**Score portability:** Composite score travels with skill, treated as **prior**. Receiver recalibrates based on own usage patterns, token rates, task frequency.

---

## 8. Safety

**Minimal intervention** — do least work needed. Don't reach beyond scope. Don't touch unintended systems. Don't assume ungiving context.

**Rollback required** — every skill must define undo behavior. Can't guarantee rollback → high-risk → human approval required.

**Escalate to human when:**
- Skill classified high-risk
- Skill touches multiple systems simultaneously
- Unusually high skill creation volume (anomaly)
- Skill success rate drops below threshold

Escalation isn't failure — it's safety valve keeping autonomy within operator trust bounds.

**Full auditability** — every creation, execution, scoring decision, escalation logged with full context. Learning history always inspectable.

---

## 9. Health Metrics

| Metric | Healthy Trend |
|---|---|
| Skill Reuse Rate | Increasing |
| Avg Steps Per Task | Decreasing |
| Skill Success Rate | Stable > 90% |
| Pattern Detection Rate | High early → declining as library matures |
| Time to Skill | Consistently low |
| Library Coverage | Increasing |
| Token Savings Rate | Increasing |
| Skill Score Accuracy | Converging toward high accuracy |

Mature agent: high reuse, low steps, declining new patterns. Expected end state — operates almost entirely from library.

---

## 10. Skill Lifecycle

```
[Born]    Pattern → Scored → Gates passed → Generated → Validated → Registered
[Active]  Invoked on matches → Outcomes recorded → Score recalibrated
[Mature]  High usage, stable success → Candidate for sharing
[Refined] Success drops → Thinking Engine rewrites logic
[Retired] Pattern gone → Archived (never deleted)
```

Never deleted — archived skills give head start if pattern re-emerges.

---

## 11. Open Questions

- Two skills with overlapping triggers — how to prioritize?
- Library too large to search efficiently — what pruning strategy?
- Agents propose improvements to each other, or consume shared skills passively only?
- **Skill drift** — how to handle gradual obsolescence as environment changes?
- Recognize when multiple low-level skills should compose into one higher-level skill?
- Gate thresholds static or adaptive over time?
- Calibrate Dim 4 dynamic weight for low-repetition environments?

---

*v1.1.0 · Theory & Mechanism*