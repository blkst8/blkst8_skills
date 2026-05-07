# Node.js Memory Leak — Causes & Debug

## Common Causes

**1. Unbounded caches / global accumulation**
Map/object grows forever, never trimmed. Each request adds entry, nothing evicts. Heap climbs linearly with traffic.

**2. Event emitter listeners never removed**
`emitter.on(...)` inside loop or request handler. Each call stacks another listener. `EventEmitter` memory leak warning fires at >10 — don't ignore it.

**3. Closures holding refs**
Closure captures large obj/buffer. Outer fn returns, but closure keeps ref alive. GC can't collect. Common in async callbacks, timers, promises.

**4. Timers / intervals not cleared**
`setInterval` inside per-request code. Interval holds closure → closure holds data → data never freed. Same with `setTimeout` chains that reschedule themselves.

**5. Streams not consumed / destroyed**
Readable stream created, never `.resume()` or `.destroy()`. Data buffers in memory waiting for consumer that never comes.

**6. Forgotten promise chains**
Rejected promise with no `.catch()`. In older Node versions, rejection hangs in memory. Even in newer — long promise chains holding large payloads.

**7. Module-level arrays / sets as queues**
Push to module-level array as "queue," never drain. Grows unbounded. Classic in home-rolled job queues or log buffers.

**8. Native addons / external memory**
`Buffer`, `TypedArray`, native addon alloc memory outside V8 heap. `process.memoryUsage().external` shows this. Heap profiler won't catch it.

---

## Debug Steps

### Step 1 — Confirm it's a leak, not normal growth
```bash
node --max-old-space-size=512 server.js
```
Watch `rss` and `heapUsed` via:
```js
setInterval(() => console.log(process.memoryUsage()), 30_000);
```
Steady climb = leak. Plateau then GC = not leak.

### Step 2 — Heap snapshot diff (V8 inspector)
Start with `--inspect` flag:
```bash
node --inspect server.js
```
Open `chrome://inspect` → Memory tab → Take snapshot. Generate load. Take second snapshot. **Select "Comparison" view** — sort by `# New` column. Biggest growers = suspects.

### Step 3 — Heap profiler in production (clinic.js)
```bash
npx clinic heapprofiler -- node server.js
```
Run load, Ctrl+C. Opens flamegraph showing allocation sites. Fastest way to find *where* memory allocated.

### Step 4 — `--expose-gc` + manual GC to isolate
```bash
node --expose-gc server.js
```
```js
global.gc();
const before = process.memoryUsage().heapUsed;
// ... run suspect code ...
global.gc();
const after = process.memoryUsage().heapUsed;
console.log('delta', after - before);
```
Positive delta after GC = retained memory = leak.

### Step 5 — Check event listeners
```js
process.on('warning', w => console.warn(w.name, w.message));
```
`MaxListenersExceededWarning` = listener leak. Track with:
```js
emitter.listenerCount('event');
```

### Step 6 — `memwatch-next` or `@nicolo-ribaudo/gc-stats`
```bash
npm i memwatch-next
```
```js
const memwatch = require('memwatch-next');
memwatch.on('leak', info => console.log('LEAK', info));
memwatch.on('stats', stats => console.log(stats));
```
Fires `leak` event when heap grows across 5+ consecutive GCs.

### Step 7 — Check native / external memory
```js
const { external, arrayBuffers } = process.memoryUsage();
```
Growing `external` or `arrayBuffers` = native addon or `Buffer` leak. `Buffer.allocUnsafe` without release, or pooled buffers not returned.

---

## Quick Checklist

- `Map`/`Set`/array at module level? Add max size + eviction.
- `emitter.on` inside request handler? Move outside or call `removeListener` / `once`.
- `setInterval` / `setTimeout` inside handler? Store ref, call `clearInterval` on cleanup.
- Streams? Always `.destroy()` on error, `.resume()` if data not needed.
- Promises? Every chain needs `.catch()`.
- Caching lib (e.g. `node-cache`, `lru-cache`)? Set `maxSize` + `ttl`.

---

## Tools Summary

| Tool | Best for |
|---|---|
| V8 heap snapshot diff | Find retained object types |
| `clinic heapprofiler` | Find allocation call sites |
| `memwatch-next` | Production leak alerts |
| `--expose-gc` delta test | Isolate single code path |
| `process.memoryUsage()` | Confirm trend / external memory |
