## `Promise.all` vs `Promise.allSettled`

Both methods accept an iterable of promises, but they differ in how they handle rejection.

### `Promise.all`

Resolves when **all** promises resolve. **Rejects immediately** if **any** promise rejects (short-circuits).

```js
const results = await Promise.all([
  fetch('/api/users'),
  fetch('/api/posts'),
  fetch('/api/comments'),
]);
// If /api/posts rejects, the whole thing rejects immediately
// The other promises are still running but their results are ignored
```

- **Resolves to**: an array of resolved values (in input order)
- **Rejects with**: the reason of the first rejection
- **Use when**: all results are required and a single failure means you can't proceed

### `Promise.allSettled`

Waits for **all** promises to finish — resolved or rejected — and **never rejects**.

```js
const results = await Promise.allSettled([
  fetch('/api/users'),
  fetch('/api/posts'),
  fetch('/api/comments'),
]);

for (const result of results) {
  if (result.status === 'fulfilled') {
    console.log(result.value);
  } else {
    console.error(result.reason);
  }
}
```

- **Resolves to**: an array of outcome objects, each with:
  - `{ status: 'fulfilled', value: ... }` or
  - `{ status: 'rejected', reason: ... }`
- **Never rejects**
- **Use when**: you want results from all promises regardless of individual failures — e.g. a dashboard loading multiple independent widgets

### Key Difference Summary

| | `Promise.all` | `Promise.allSettled` |
|---|---|---|
| Short-circuits on rejection | Yes | No |
| Can reject | Yes | No |
| Result shape | Array of values | Array of `{status, value/reason}` |
| Use case | All-or-nothing | Best-effort, partial success OK |

### Quick Rule of Thumb

- Need **every** result to succeed → `Promise.all`
- Want to **inspect** what succeeded and what failed → `Promise.allSettled`
