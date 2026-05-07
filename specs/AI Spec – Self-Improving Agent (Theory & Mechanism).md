# AI Specification: Self-Improving Agent with Auto-Skill Creation
> **Version:** 1.1.0
> **Date:** 2026-05-07
> **Status:** Draft

---

## 1. Overview

### 1.1 Purpose
This document defines the **theory, mechanism, and behavioral specification** of a Self-Improving AI Agent — a system that continuously observes its own task execution, discovers recurring patterns through internal reasoning, and autonomously creates reusable skills from those patterns to improve its future performance.

### 1.2 Core Philosophy

> *An agent should never solve the same problem twice the same way.*

Traditional AI agents are static — they execute tasks using pre-defined tools and stop there. A Self-Improving Agent treats **every task execution as a learning opportunity**. It watches itself work, thinks critically about what it did, and permanently encodes better ways of doing things into its own skill set.

This is not fine-tuning. This is not retraining. This is **runtime self-improvement** — the agent grows smarter within its operational lifetime, without any human intervention.

### 1.3 Always-On Principle
Self-improvement is **not an optional mode**. It is a fundamental, always-enabled layer of the agent's cognition. Every task the agent runs — regardless of success or failure — passes through the reflection and pattern-detection pipeline. The agent is always watching itself.

---

## 2. Theoretical Foundation

### 2.1 The Three Pillars

The system is built on three cognitive principles borrowed from human learning theory:

**Pillar 1 — Metacognition (Thinking About Thinking)**
The agent maintains a continuous inner monologue about its own performance. After every task, it does not simply return a result — it asks: *"How did I do that? Could I have done it better? Have I done something like this before?"* This self-awareness is the seed of all improvement.

**Pillar 2 — Pattern Abstraction**
Humans learn by abstracting specific experiences into general rules. When a person burns their hand on a stove, they don't just remember *"that specific stove is hot"* — they abstract the rule *"hot surfaces are dangerous."* The agent operates the same way: it abstracts specific task struggles into general, reusable patterns.

**Pillar 3 — Skill Encoding**
Once a pattern is abstracted, it must be encoded into a durable, executable form — a **skill**. A skill is the agent's way of crystallizing a lesson into action. Skills are not descriptions; they are capabilities. They can be invoked, measured, improved, and shared.

### 2.2 The Improvement Feedback Loop

Self-improvement is a closed loop, not a one-time event:

```
Experience → Reflection → Pattern → Skill Scoring → Skill → Better Experience → Deeper Reflection → ...
```

Each cycle of this loop produces an agent that is measurably more capable than the previous cycle. Over time, the agent builds a rich internal library of skills that covers more and more of its operational domain — requiring less and less manual effort per task.

### 2.3 Why This Works at Runtime

Classical machine learning improves models through gradient descent on large datasets — a slow, expensive, offline process. This system achieves improvement through a different mechanism: **symbolic skill accumulation**.

Instead of changing the agent's weights, we change the agent's **available actions**. Each new skill expands what the agent can do in a single step. The agent's intelligence grows not by becoming a different model, but by having access to an ever-richer vocabulary of capabilities.

---

## 3. Core Concepts

| Concept | Definition |
|---|---|
| **Task** | Any unit of work the agent is asked to perform |
| **Execution Trace** | The full record of every step, decision, and outcome during a task |
| **Reflection** | The agent's structured reasoning about its own execution trace |
| **Pattern** | A recurring, nameable behavior observed across one or more execution traces |
| **Skill** | A named, reusable capability derived from a pattern — permanently stored and invocable |
| **Skill Score** | A composite rating that determines whether a skill is worth creating and how urgently |
| **Skill Library** | The agent's persistent, growing collection of all created skills |
| **Thinking Engine** | The reasoning module (powered by LLM) responsible for reflection and skill design |
| **Pattern Fingerprint** | A unique semantic identifier for a pattern, used to match future occurrences |

---

## 4. The Self-Improvement Mechanism

The mechanism operates as a **six-phase cycle** that runs after every task execution.

---

### Phase 1 — Observe (Execution Tracing)

As the agent executes a task, every action it takes is silently recorded into an **Execution Trace**. This trace is a complete, ordered log of:

- What steps were taken and in what order
- What decisions were made at each branch point
- What inputs were consumed and what outputs were produced
- Where errors occurred and how they were handled
- How long each step took
- Which steps were repeated
- How many tokens were consumed during execution

The agent does not modify its behavior during execution to accommodate tracing — it simply acts, and the trace is collected passively in the background. This ensures the agent's performance is never degraded by the observation process.

**Key insight:** The trace is not just a log for debugging. It is the raw material from which all learning originates.

---

### Phase 2 — Reflect (Internal Reasoning)

Once a task is complete, the agent enters a **Reflection Phase**. This is where the agent's Thinking Engine examines the execution trace and asks a structured set of introspective questions:

- *Did I repeat any action more than once?*
- *Did I take more steps than should have been necessary?*
- *Did I encounter an error and have to recover?*
- *Was there a sequence of steps that always appeared together?*
- *Did I process data in a way that felt mechanical and predictable?*
- *How many tokens did I consume, and was that proportionate to the task's complexity?*
- *If this task repeats in the future, how expensive will it be without a skill?*

Reflection is not a simple rule check. It is a **reasoning process** — the Thinking Engine reads the trace the way a thoughtful person reads a journal entry, looking for themes, inefficiencies, and moments of struggle.

The output of reflection is a **Reflection Report**: a structured assessment of whether a meaningful pattern exists in the trace, what that pattern is, and how confident the agent is in its assessment.

**Reflection is always triggered.** Whether the task succeeded or failed, whether it was simple or complex — the agent always reflects. A successful task might reveal an efficiency pattern. A failed task might reveal an error-handling pattern. Both are valuable.

---

### Phase 3 — Extract (Pattern Identification)

If the Reflection Report identifies a pattern, the agent moves to **Pattern Extraction**. This phase transforms the raw observation from reflection into a clean, normalized, reusable pattern definition.

Pattern extraction involves three operations:

**Naming:** The agent assigns a concise, human-readable name to the pattern. The name must be general enough to apply to future occurrences, not specific to the current task. For example: *"retry on transient failure"* rather than *"retry the weather API call."*

**Classifying:** The pattern is assigned a type from a defined taxonomy:

| Pattern Type | Description |
|---|---|
| **Repetition** | The same action was performed multiple times unnecessarily |
| **Inefficiency** | The task required far more steps than the problem warranted |
| **Error Recovery** | The agent encountered a failure and had to recover manually |
| **Data Transformation** | The same shape of data was processed in the same way repeatedly |
| **Decision Sequence** | A predictable chain of decisions always led to the same outcome |
| **Resource Management** | Acquisition and release of a resource followed a consistent pattern |

**Fingerprinting:** The pattern is given a unique semantic fingerprint — a vector embedding of its normalized description. This fingerprint is used in the future to recognize when the same pattern appears again, even if the surface-level task looks different.

---

### Phase 4 — Score (Skill Worthiness Evaluation)

Before the agent invests in generating a skill, it must answer a critical question:

> *Is this pattern worth turning into a skill right now?*

Not every pattern justifies the cost of skill creation. Some patterns are rare, cheap to re-execute, or too specific to generalize. The **Skill Scoring System** is the mechanism by which the agent makes this judgment objectively and consistently.

The agent evaluates the candidate skill against **four scoring dimensions**. Each dimension produces a score, and the scores are combined into a **Composite Skill Score** that determines whether skill creation proceeds.

If the Composite Score meets the creation threshold, skill generation begins. If it does not, the pattern is noted in memory but no skill is created — the agent will re-evaluate if the pattern appears again.

---

#### 4.1 Scoring Dimensions

---

##### Dimension 1 — Speed Gain
**"How much faster will future tasks be if this skill exists?"**

This dimension measures the expected reduction in execution time when the skill handles this pattern instead of the agent solving it manually. It accounts for the elimination of redundant reasoning steps, the removal of trial-and-error cycles, and the replacement of multi-step sequences with a single direct invocation.

**Scoring rule:** This score must always be greater than **6 out of 10**.
A skill that does not meaningfully accelerate execution is not worth creating. Speed gain is a hard gate — if a candidate skill cannot demonstrate substantial time savings, it is rejected regardless of its other scores.

| Score | Meaning |
|---|---|
| 1 – 3 | Negligible speed improvement — task is already fast |
| 4 – 6 | Moderate improvement — borderline, fails the gate |
| 7 – 8 | Significant improvement — task becomes noticeably faster |
| 9 – 10 | Dramatic improvement — task collapses from many steps to one |

---

##### Dimension 2 — Token Efficiency
**"How token-expensive is this pattern to handle manually, and how much does the skill reduce that cost?"**

Every time the agent reasons through a pattern without a skill, it consumes tokens — for thinking, for planning, for generating intermediate outputs. This dimension measures the token cost of the current manual approach and scores how efficiently the skill would replace it.

**Scoring rule:** This score must be lower than **3 out of 10**.
A low score here is desirable — it means the skill dramatically reduces token consumption. A high score means the manual approach is already token-efficient, and a skill would offer little savings. Skills that score 3 or above on this dimension are deprioritized, as their token-saving benefit is insufficient to justify creation.

| Score | Meaning |
|---|---|
| 1 – 2 | Skill eliminates nearly all token overhead — high savings |
| 3 | Borderline — skill offers modest savings, fails the gate |
| 4 – 6 | Manual approach is moderately efficient — low benefit |
| 7 – 10 | Manual approach is already very token-lean — skill adds little value |

---

##### Dimension 3 — Repetition Density
**"How often does this pattern appear, and how concentrated is that repetition?"**

A pattern that appears once is an anecdote. A pattern that appears repeatedly is a law. This dimension measures how frequently the pattern has been observed — both within a single task (intra-task repetition) and across multiple tasks over time (inter-task repetition).

**Scoring rule:** This score must be greater than **4 out of 10**.
A skill built for a pattern that rarely repeats will sit unused in the library. Repetition density ensures the agent only creates skills for patterns that have demonstrated a meaningful frequency of occurrence.

| Score | Meaning |
|---|---|
| 1 – 2 | Pattern seen once or twice — too rare to justify a skill |
| 3 – 4 | Pattern seen occasionally — borderline, fails the gate |
| 5 – 6 | Pattern repeats regularly — skill will get meaningful use |
| 7 – 10 | Pattern is extremely frequent — skill is urgently needed |

---

##### Dimension 4 — Future Token Cost (Weighted)
**"If this task repeats in the future without a skill, how expensive will that be in total?"**

This dimension looks forward, not backward. It estimates the cumulative token cost of *not* creating the skill — the total tokens that will be spent on this pattern across all future occurrences if it remains unautomated.

**This dimension is dynamically weighted.** If the projected future token cost is high — meaning the pattern is both token-intensive and likely to repeat many times — the weight of this dimension in the Composite Score increases automatically. The more expensive the future is without the skill, the more urgently the skill should be created.

Conversely, if the future token cost is low — because the pattern is rare or already cheap to handle — this dimension contributes less to the final score, and other dimensions carry more influence.

| Future Token Cost | Weight Applied |
|---|---|
| Very High (pattern is frequent + token-heavy) | Maximum weight — this dimension dominates the score |
| High | Elevated weight — strong influence on composite |
| Moderate | Standard weight — equal influence with other dimensions |
| Low | Reduced weight — other dimensions take precedence |
| Very Low | Minimal weight — near-irrelevant to the decision |

---

#### 4.2 Composite Skill Score

The four dimension scores are combined into a single **Composite Skill Score** using a weighted formula. The weights are not fixed — Dimension 4's weight adjusts dynamically based on projected future token cost, as described above.

The composite score is evaluated against the following outcome table:

| Composite Score | Decision |
|---|---|
| All gates passed + high composite | **Create skill immediately** |
| All gates passed + moderate composite | **Create skill, lower priority** |
| One gate failed | **Defer — re-evaluate after next occurrence** |
| Two or more gates failed | **Reject — pattern does not warrant a skill** |
| Speed gate failed (score ≤ 6) | **Hard reject — no skill created regardless of other scores** |

**Gate Summary:**

| Dimension | Gate Condition | Gate Type |
|---|---|---|
| Speed Gain | Score **> 6** | Hard gate — rejection if failed |
| Token Efficiency | Score **< 3** | Soft gate — deprioritizes if failed |
| Repetition Density | Score **> 4** | Soft gate — defers if failed |
| Future Token Cost | No fixed gate — dynamic weight | Weight modifier only |

---

#### 4.3 Scoring Rationale

The asymmetry in the gate conditions is intentional and reflects the relative importance of each dimension:

- **Speed is non-negotiable.** A skill that doesn't make the agent meaningfully faster defeats the entire purpose of self-improvement. Hence the hard gate.
- **Token efficiency is a strong signal but not absolute.** A skill might be worth creating even if token savings are modest, if repetition density or future cost is very high.
- **Repetition is a prerequisite for value.** A skill for a one-time pattern is waste. The soft gate ensures skills are only created for patterns with demonstrated recurrence.
- **Future cost is the most forward-looking signal.** It captures the compounding cost of inaction. Its dynamic weight ensures it has maximum influence precisely when it matters most — when the cost of not acting is highest.

---

### Phase 5 — Generate (Skill Creation)

Once a pattern passes the Skill Scoring phase, the Thinking Engine designs a **Skill** to handle it.

Skill generation is a creative act of reasoning, not template filling. The Thinking Engine must:

1. **Understand the pattern deeply** — What is the essence of what keeps happening? What is the agent really trying to accomplish when this pattern appears?

2. **Design the skill's interface** — What information does the skill need as input? What should it produce as output? What are the preconditions for the skill to be applicable?

3. **Define the skill's behavior** — What sequence of actions should the skill perform? How should it handle edge cases? What does success look like? What does failure look like, and how should it be handled?

4. **Specify rollback behavior** — If the skill's actions cannot be completed, how should the agent undo any partial changes to restore a clean state?

5. **Write test cases** — What are the simplest examples that prove the skill works correctly?

The result is a **Skill Blueprint**: a complete specification of a new capability, ready to be validated and registered.

**Critical principle:** Skills must be *general*, not *specific*. A skill generated from a task about retrying an API call must be general enough to handle retrying a database query, a file read, or any other transient failure — not just the original API.

---

### Phase 6 — Register (Skill Persistence)

Before a skill enters the library, it passes through **validation**:

- **Safety Review:** Does the skill attempt any dangerous or irreversible operations? Does it have appropriate error handling? Is its blast radius — the scope of systems it can affect — appropriately limited?
- **Correctness Test:** Does the skill produce the correct output when given the test cases defined during generation?
- **Risk Classification:** The skill is assigned a risk level (Low, Medium, High) based on the reversibility of its actions and the breadth of systems it touches.

Skills that pass validation are **registered in the Skill Library** with full metadata: their name, pattern trigger, fingerprint, risk level, Composite Skill Score, creation timestamp, and performance statistics.

From this moment forward, whenever the agent encounters a task and finds a pattern that matches a registered skill's fingerprint, it invokes the skill directly — bypassing the need to solve the problem from scratch.

---

## 5. The Skill Library

### 5.1 Nature of the Library

The Skill Library is the agent's **long-term procedural memory**. It is the tangible product of all self-improvement cycles. Unlike the agent's context window — which is temporary and task-scoped — the Skill Library persists indefinitely across all tasks and sessions.

The library grows over time. Every pattern the agent encounters that does not yet have a skill becomes a new entry. The agent never forgets a skill it has learned.

### 5.2 Skill Retrieval

When a new task begins, the agent first consults the Skill Library before taking any manual action. Retrieval works in two modes:

**Exact Match:** If the task description or context contains a phrase that directly matches a skill's pattern trigger, that skill is immediately selected.

**Semantic Match:** If no exact match exists, the agent generates an embedding of the current task context and performs a similarity search across all skill fingerprints. If a skill is found with similarity above a defined threshold, it is selected — even if the surface-level wording is completely different.

If no match is found at all, the agent proceeds manually and the self-improvement cycle runs after completion.

### 5.3 Skill Evolution

Skills are not static. Every time a skill is executed, its outcome is recorded. Over time, the agent can observe:

- Skills with declining success rates may need refinement
- Skills that are never invoked may be obsolete
- Two skills that are frequently used together may be candidates for merging into a single, more powerful skill
- A skill's Composite Score is recalculated periodically as new usage data accumulates, ensuring scores remain accurate over time

This creates a **second-order improvement loop**: not just creating new skills, but improving existing ones.

---

## 6. Always-On Behavior

### 6.1 Why Self-Improvement Must Always Be Active

Self-improvement is not a mode the agent enters under special conditions. It is a **continuous cognitive layer** that runs alongside every task, always. This is a deliberate design decision rooted in three principles:

**Consistency:** If self-improvement only activates sometimes, the agent's behavior becomes unpredictable. Operators cannot reason about when the agent is learning and when it is not.

**Completeness:** Patterns do not announce themselves. A pattern that seems trivial on first occurrence may become critically important on the hundredth. The only way to catch all patterns is to observe all executions.

**Compounding Returns:** The value of self-improvement compounds. An agent that reflects on every task builds a richer, more nuanced skill library than one that reflects occasionally. The difference in capability between the two agents grows exponentially over time.

### 6.2 Performance Overhead

The always-on nature of self-improvement is designed to have **negligible impact on task performance**. Reflection, pattern extraction, and skill scoring happen *after* the task is complete, not during it. Skill retrieval at the start of a task is a fast lookup operation. The agent's primary task execution is never blocked or slowed by the improvement pipeline.

---

## 7. Multi-Agent Skill Sharing

### 7.1 The Collective Intelligence Model

When multiple agents operate in the same environment, the Skill Library can be shared across all of them. This transforms individual learning into **collective intelligence**: a skill discovered by one agent immediately becomes available to all others.

The mechanism is straightforward: when Agent A creates a skill and marks it as shared, it enters a common pool. When Agent B encounters a pattern, it searches not only its own skills but the shared pool as well. If Agent A's skill matches, Agent B uses it — without needing to rediscover the pattern or regenerate the skill.

### 7.2 Trust and Propagation Rules

Not all skills are shared automatically. Sharing follows a trust model based on risk level:

- **Low-risk skills** are shared automatically and immediately
- **Medium-risk skills** require acknowledgment from the originating agent before sharing
- **High-risk skills** are never shared and remain isolated to the agent that created them

This ensures that the collective skill pool remains safe and that no single agent can introduce dangerous capabilities into the shared environment.

### 7.3 Score Portability

When a skill is shared between agents, its Composite Skill Score travels with it — but is treated as a **prior**, not a fact. The receiving agent recalibrates the score based on its own usage patterns, token consumption rates, and task frequency. Over time, each agent holds a version of the score that reflects its own operational reality, while still benefiting from the originating agent's initial signal.

---

## 8. Safety Model

### 8.1 The Principle of Minimal Intervention

Every skill the agent creates must adhere to the **Principle of Minimal Intervention**: a skill should do the least amount of work necessary to handle its pattern. It should not reach beyond its defined scope, modify systems it was not designed to touch, or make assumptions about context it was not given.

### 8.2 Reversibility Requirement

All skills must define a rollback behavior. If a skill begins executing and encounters an unrecoverable error midway through, it must be able to undo any changes it has already made. A skill that cannot guarantee rollback is classified as high-risk and requires human approval before registration.

### 8.3 Human Escalation

The agent escalates to a human operator in the following situations:

- A generated skill is classified as high-risk
- A skill would affect more than one system or service simultaneously
- The agent generates an unusually high number of skills in a short period (anomaly detection)
- A skill's success rate falls below an acceptable threshold after repeated use

Human escalation is not a failure state — it is a designed safety valve that ensures the agent's autonomy never exceeds the trust boundaries set by its operators.

### 8.4 Full Auditability

Every skill creation event, every skill execution, every scoring decision, and every escalation is permanently logged with full context: which agent triggered it, what pattern was detected, what scores were assigned and why, what skill was created or used, what the outcome was, and when it happened. The agent's entire learning history is transparent and inspectable at any time.

---

## 9. Improvement Metrics

The health and progress of the self-improvement system is measured by the following indicators:

| Metric | What It Measures | Healthy Trend |
|---|---|---|
| **Skill Reuse Rate** | Percentage of tasks handled by an existing skill | Increasing over time |
| **Average Steps Per Task** | How many manual steps the agent takes per task | Decreasing over time |
| **Skill Success Rate** | Percentage of skill executions that succeed | Stable above 90% |
| **Pattern Detection Rate** | Percentage of tasks where a new pattern is found | High early, declining as library matures |
| **Time to Skill** | Time elapsed from pattern detection to skill registration | Consistently low |
| **Library Coverage** | Percentage of the agent's task domain covered by skills | Increasing over time |
| **Token Savings Rate** | Tokens saved per task due to skill reuse vs. manual execution | Increasing over time |
| **Skill Score Accuracy** | How well Composite Scores predicted actual skill value post-creation | Converging toward high accuracy |

A mature agent — one that has been running for a long time in a stable domain — will show a high skill reuse rate, low average steps per task, and a declining pattern detection rate. This is the expected and desired end state: the agent has learned most of what there is to learn in its domain and now operates almost entirely from its skill library.

---

## 10. Lifecycle of a Skill

```
[Born]      Pattern detected → Scored → Gates passed → Skill generated → Validated → Registered
    │
[Active]    Invoked on matching tasks → Success/failure recorded → Score recalculated
    │
[Mature]    High usage, stable success rate, score confirmed → Candidate for sharing
    │
[Refined]   Success rate drops or score degrades → Thinking Engine rewrites skill logic
    │
[Retired]   Pattern no longer appears → Skill archived (never deleted)
```

Skills are never permanently deleted. Retired skills are archived so that if a pattern re-emerges in the future, the agent has a starting point for the new skill rather than starting from scratch.

---

## 11. Open Design Questions

- How should two skills with overlapping pattern triggers be prioritized against each other?
- At what point does a skill library become too large to search efficiently, and what pruning strategy should be applied?
- Should agents be able to propose skill improvements to each other, or only consume shared skills passively?
- How should the system handle **skill drift** — the gradual obsolescence of skills as the environment they were designed for changes over time?
- Can the agent recognize when multiple low-level skills should be composed into a single higher-level skill?
- Should Composite Score gate thresholds be static, or should they adapt over time as the agent learns more about its own domain?
- How should the dynamic weight of Dimension 4 (Future Token Cost) be calibrated for agents operating in low-repetition environments?

---

*Self-Improving Agent System Specification — v1.1.0*
*Theory & Mechanism Edition*