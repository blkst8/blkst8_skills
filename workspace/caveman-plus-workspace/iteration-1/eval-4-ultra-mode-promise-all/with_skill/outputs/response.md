`Promise.all` vs `Promise.allSettled`:

**`Promise.all`** — fail-fast. One rejects → whole rejects immediately. Returns resolved-values array.

**`Promise.allSettled`** — waits all settle. Never rejects. Returns `[{status, value}]` or `[{status, reason}]` per promise.

Use `all` → need all succeed, bail on first fail.
Use `allSettled` → need all results regardless of failures.
