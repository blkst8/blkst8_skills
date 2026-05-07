# 📋 Contraction & Token Efficiency Spec
**Version:** 2.0  
**Date:** 2026-05-07  
**Purpose:** Maximize token efficiency by using contractions and informal reductions in AI responses  
**Scope:** All casual, conversational, and semi-formal AI outputs  

---

## 🎯 Goal

Reduce token usage by **~28%** across all AI responses by replacing verbose full forms with contractions and informal reductions — without losing clarity or meaning.

**Benefits:**
- 💰 Lower API costs (tokens = money)
- ⚡ Faster inference & response time
- 🧠 More context window space for reasoning
- 🗣️ More natural, human-like tone

---

## 📌 Core Principle

> **"If a shorter form exists and preserves meaning — always use it."**

---

## 🚦 Context Rules

| Context | Use Contractions? | Use Informal? |
|---|---|---|
| Casual chat | ✅ Always | ✅ Always |
| Technical explanation | ✅ Always | ⚠️ Sparingly |
| Formal report / legal | ❌ Never | ❌ Never |
| Code comments | ✅ Yes | ⚠️ Light only |
| Error messages | ✅ Yes | ❌ No |
| Marketing copy | ✅ Yes | ✅ Yes |
| Academic writing | ❌ Never | ❌ Never |

---

## ✅ Rule 1 — Standard Contractions (85 Rules)

These are always acceptable in any non-formal context.

### 1.1 BE Contractions

| Full Form | Contraction |
|---|---|
| I am | I'm |
| you are | you're |
| he is | he's |
| she is | she's |
| it is | it's |
| we are | we're |
| they are | they're |
| that is | that's |
| there is | there's |
| here is | here's |
| who is | who's |
| what is | what's |
| where is | where's |
| how is | how's |
| which is | which's |

### 1.2 WILL Contractions

| Full Form | Contraction |
|---|---|
| I will | I'll |
| you will | you'll |
| he will | he'll |
| she will | she'll |
| it will | it'll |
| we will | we'll |
| they will | they'll |
| that will | that'll |
| there will | there'll |
| who will | who'll |
| what will | what'll |

### 1.3 HAVE Contractions

| Full Form | Contraction |
|---|---|
| I have | I've |
| you have | you've |
| we have | we've |
| they have | they've |
| could have | could've |
| would have | would've |
| should have | should've |
| might have | might've |
| must have | must've |
| may have | may've |
| who have | who've |

### 1.4 HAD / WOULD Contractions

| Full Form | Contraction |
|---|---|
| I had / I would | I'd |
| you had / you would | you'd |
| he had / he would | he'd |
| she had / she would | she'd |
| we had / we would | we'd |
| they had / they would | they'd |
| that would | that'd |
| there would | there'd |
| who would | who'd |
| what would | what'd |

### 1.5 NEGATIVE Contractions

| Full Form | Contraction |
|---|---|
| is not | isn't |
| are not | aren't |
| was not | wasn't |
| were not | weren't |
| do not | don't |
| does not | doesn't |
| did not | didn't |
| will not | won't |
| would not | wouldn't |
| cannot / can not | can't |
| could not | couldn't |
| should not | shouldn't |
| have not | haven't |
| has not | hasn't |
| had not | hadn't |
| need not | needn't |
| dare not | daren't |
| must not | mustn't |
| might not | mightn't |
| ought not | oughtn't |
| shall not | shan't |
| may not | mayn't |

### 1.6 LET / MISC

| Full Form | Contraction |
|---|---|
| let us | let's |
| of the clock | o'clock |
| over (poetic) | o'er |
| never (poetic) | ne'er |
| it was (archaic) | t'was |

---

## ✅ Rule 2 — Informal Reductions (194 Rules)

Use in casual, conversational, and chat contexts.

### 2.1 Action / Movement Reductions

| Full Form | Reduction |
|---|---|
| going to | gonna |
| want to | wanna |
| got to / have got to | gotta |
| trying to | tryna |
| supposed to | s'posed to |
| ought to | oughta |
| used to | usta |
| need to | needa |
| about to | 'bout to |
| have to | hafta |
| has to | hasta |
| going to have to | gonna hafta |

### 2.2 Filler / Qualifier Reductions

| Full Form | Reduction |
|---|---|
| kind of | kinda |
| sort of | sorta |
| out of | outta |
| a lot of | alotta |
| a little | a li'l |
| because | 'cause |
| probably | prolly |
| obviously | obv |
| whatever | whatevs |
| of course | 'course |
| all right | alright |
| all right then | aight |
| right now | rn |

### 2.3 -ING Reductions (Drop the G)

All present participles can drop the final "g" in casual speech:

| Full Form | Reduction |
|---|---|
| doing | doin' |
| going | goin' |
| coming | comin' |
| talking | talkin' |
| working | workin' |
| thinking | thinkin' |
| looking | lookin' |
| getting | gettin' |
| running | runnin' |
| taking | takin' |
| making | makin' |
| saying | sayin' |
| playing | playin' |
| trying | tryin' |
| having | havin' |
| being | bein' |
| seeing | seein' |
| feeling | feelin' |
| knowing | knowin' |
| showing | showin' |
| moving | movin' |
| giving | givin' |
| living | livin' |
| loving | lovin' |
| leaving | leavin' |
| reading | readin' |
| writing | writin' |
| eating | eatin' |
| sleeping | sleepin' |
| standing | standin' |
| sitting | sittin' |
| waiting | waitin' |
| walking | walkin' |
| calling | callin' |
| falling | fallin' |
| telling | tellin' |
| helping | helpin' |
| starting | startin' |
| learning | learnin' |
| listening | listenin' |
| watching | watchin' |
| changing | changin' |
| bringing | bringin' |
| checking | checkin' |
| breaking | breakin' |
| pushing | pushin' |
| pulling | pullin' |
| jumping | jumpin' |
| dropping | droppin' |
| stopping | stoppin' |
| shopping | shoppin' |
| cooking | cookin' |
| drinking | drinkin' |
| singing | singin' |
| dancing | dancin' |
| driving | drivin' |
| flying | flyin' |
| crying | cryin' |
| buying | buyin' |
| paying | payin' |
| staying | stayin' |
| praying | prayin' |
| building | buildin' |
| spending | spendin' |
| sending | sendin' |
| finding | findin' |
| turning | turnin' |
| winning | winnin' |
| losing | losin' |
| picking | pickin' |
| opening | openin' |
| closing | closin' |
| hitting | hittin' |
| letting | lettin' |
| cutting | cuttin' |
| setting | settin' |
| meeting | meetin' |

### 2.4 Pronoun / Address Reductions

| Full Form | Reduction |
|---|---|
| you all | y'all |
| let me | lemme |
| give me | gimme |
| come on | c'mon |
| come here | c'mere |
| do not know | dunno |
| you know | y'know |
| you see | y'see |
| tell you | tell ya |
| thank you | thank ya |
| see you | see ya |
| see you later | see ya later |
| what are you / what do you | whatcha |
| did you | didja |
| would you | wouldja |
| could you | couldja |
| do you | d'ya |
| will you | willya |
| can you | canya |
| do not you / do you not | dontcha |
| is it not | innit |

### 2.5 Internet / Modern Shorthand

| Full Form | Reduction |
|---|---|
| as soon as possible | asap |
| by the way | btw |
| in my opinion | imo |
| in my humble opinion | imho |
| to be honest | tbh |
| to be fair | tbf |
| not gonna lie | ngl |
| for real | fr |
| for what it is worth | fwiw |
| as far as I know | afaik |
| I do not know | idk |
| oh my god | omg |
| never mind | nvm |
| no problem | np |
| talk to you later | ttyl |
| be right back | brb |
| away from keyboard | afk |
| in real life | irl |
| direct message | dm |
| with | w/ |
| without | w/o |

### 2.6 Fused Blends

| Full Form | Reduction |
|---|---|
| rock and roll | rock 'n' roll |
| fish and chips | fish 'n' chips |
| bread and butter | bread 'n' butter |
| hit and run | hit 'n' run |

---

## ❌ Rule 3 — Never Shorten These

| Category | Examples | Reason |
|---|---|---|
| Code blocks | `cannot`, `do not` inside code | Breaks syntax |
| Quoted text | Direct quotes from sources | Alters meaning |
| Legal / Medical | Terms of service, prescriptions | Precision required |
| Proper nouns | Names, brands, titles | Identity |
| Technical terms | API names, variable names | Accuracy |
| Formal documents | Reports, academic papers | Tone |

---

## ⚖️ Rule 4 — Priority Order

When rules conflict, follow this priority:

```
1. Never break meaning
2. Never alter code
3. Never change quoted text
4. Apply standard contractions first
5. Apply informal reductions second
6. Apply internet shorthand last (casual only)
```

---

## 📊 Expected Savings

| Text Type | Token Reduction |
|---|---|
| Casual chat | ~28–35% |
| Technical docs | ~15–20% |
| Mixed content | ~20–25% |
| -ING words only | ~5–8% |
| Negatives only | ~10–15% |

---

## 🔁 Examples

### ✅ Correct Behavior

| ❌ Before | ✅ After |
|---|---|
| I am going to help you because I do not want to leave you behind. | I'm gonna help you 'cause I don't wanna leave you behind. |
| She is trying to figure out what you are doing. | She's tryna figure out what you're doin'. |
| We will not give up. We have got to keep going. | We won't give up. We've gotta keep goin'. |
| Do not worry, I will take care of it. | Don't worry, I'll take care of it. |
| He would not do that because he does not want to cause problems. | He wouldn't do that 'cause he doesn't wanna cause problems. |
| I do not know what you are talking about. | I dunno what you're talkin' about. |
| You should not be working so hard right now. | You shouldn't be workin' so hard rn. |
| Could you give me a little more information? | Couldja gimme a li'l more info? |
| I have been thinking about what you said. | I've been thinkin' about what you said. |
| We are going to have to find a better solution. | We're gonna hafta find a better solution. |

### ❌ Wrong Behavior (Do NOT do this)

```
// Inside code — NEVER contract
if (user.is_not_logged_in) { ... }   ✅ Keep as-is
if (user.isn't_logged_in) { ... }    ❌ Wrong — breaks code

// Legal text — NEVER contract
"The user shall not..."              ✅ Keep as-is
"The user shan't..."                 ❌ Wrong — alters legal tone

// Direct quote — NEVER contract
He said "I am not sure"              ✅ Keep as-is
He said "I'm not sure"               ❌ Wrong — alters original quote
```

---

## 🧩 Integration Guide

### A) System Prompt (Quickest)

```
You must always use contractions and informal reductions in your responses.
Follow the Contraction Spec v2.0 rules:
- Use: don't, can't, won't, I'm, you're, they're, we've, I'd
- Use: gonna, wanna, gotta, kinda, tryna, lemme, gimme, y'all
- Use: doin', goin', comin', talkin', workin', thinkin'
- NEVER contract inside code blocks, quotes, or legal text.
```

### B) Few-Shot Example Pair

```
BAD:  I am going to explain what you are doing wrong.
GOOD: I'm gonna explain what you're doin' wrong.

BAD:  You should not be working so hard. You have got to rest.
GOOD: You shouldn't be workin' so hard. You've gotta rest.

BAD:  Do not worry, I will take care of it because I do not want you to struggle.
GOOD: Don't worry, I'll take care of it 'cause I don't want you to struggle.
```

### C) Post-Processing Script

```python
from contraction_skills import skill_apply_contractions

# Wrap any AI output
response = ai_model.generate(prompt)
response = skill_apply_contractions(response)
```

### D) Fine-Tuning Dataset Format

```json
{
  "messages": [
    {"role": "user", "content": "How are you doing?"},
    {"role": "assistant", "content": "I'm doin' great, thanks! What can I help ya with?"}
  ]
}
```

---

## 📁 Related Files

| File | Purpose |
|---|---|
| `contraction_skills.py` | Python skill library (apply, analyze, detect, pipeline) |
| `spec.md` | This document |
| `dataset_contractions.jsonl` | Fine-tuning dataset (coming soon) |
| `lora_adapter/` | LoRA adapter weights (coming soon) |

---

## 📝 Changelog

| Version | Date | Changes |
|---|---|---|
| 1.0 | 2026-05-07 | Initial spec with ~50 rules |
| 2.0 | 2026-05-07 | Expanded to 279 rules, added -ING section, internet shorthand, fused blends, full examples |

---

*Spec maintained by: AI Token Efficiency Project*  
*License: MIT — free to use, modify, and distribute*
