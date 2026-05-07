# Node.js Memory Leak: Common Causes and Debugging

## Most Common Causes

### 1. Event Listener Accumulation
The most frequent culprit. Every time you call `emitter.on(...)` without a corresponding `removeListener`, you add a listener that holds a reference to its closure and any captured variables. In a long-running service, attaching listeners inside request handlers or timers without cleanup is a slow, steady leak.

```js
// Bad — adds a new listener on every request
app.get('/data', (req, res) => {
  someEmitter.on('data', (chunk) => res.write(chunk)); // never removed
});
```

### 2. Unbounded Caches / Global Maps
In-memory caches that grow forever. A `Map` or plain object used as a cache with no eviction strategy will hold references to values indefinitely.

```js
const cache = new Map();
cache.set(userId, largeObject); // never evicted
```

Use an LRU cache (e.g., `lru-cache`) with a max size.

### 3. Closures Capturing Large Objects
A closure keeps the entire scope chain alive. If a closure is stored somewhere long-lived (timer, cache, global), it prevents the objects it references from being GC'd — even if you only intended to keep a small piece of data.

```js
function processRequest(largeBuffer) {
  const id = largeBuffer.id; // you only need `id`
  setTimeout(() => doWork(largeBuffer), 5000); // but largeBuffer stays alive
}
```

### 4. Timers and Intervals Never Cleared
`setInterval` callbacks hold references to their closure scope. If you create intervals inside a class or per-request scope and never call `clearInterval`, both the interval and everything it references stays alive.

### 5. Forgotten Async Operations / Pending Promises
Promises that never resolve or reject keep their closure alive. This often happens with hanging DB queries, long-polling, or incomplete stream consumers.

### 6. Streams Not Consumed or Destroyed
An unconsumed readable stream buffers data internally. If you attach a stream to a request but never pipe it or call `.destroy()` on error paths, the buffer accumulates.

### 7. Third-Party Module Leaks
Logging libraries, APM agents, or DB drivers that register global listeners or accumulate internal state. Worth checking when you've ruled out your own code.

---

## How to Debug

### Step 1 — Confirm It's a Heap Leak
First, rule out RSS growth from native modules or the OS not reclaiming memory. Use the `--expose-gc` flag and compare heap snapshots over time, not just process memory.

```bash
node --expose-gc server.js
```

Then check `process.memoryUsage()` periodically:
```js
setInterval(() => {
  const { heapUsed, heapTotal, external } = process.memoryUsage();
  console.log({ heapUsed, heapTotal, external });
}, 30_000);
```

If `heapUsed` grows consistently without plateauing, you have a heap leak.

### Step 2 — Take Heap Snapshots
Use the built-in `v8` module or Chrome DevTools Protocol to capture heap snapshots at two points in time — before the leak starts and after significant growth.

```js
const v8 = require('v8');
const fs = require('fs');

// capture a snapshot
const snapshot = v8.writeHeapSnapshot();
console.log('Snapshot written to', snapshot);
```

Or use `--inspect` and connect Chrome DevTools:
```bash
node --inspect server.js
```
Open `chrome://inspect`, go to the Memory tab, and take two snapshots separated by load. Use the **Comparison** view to see what objects are accumulating.

### Step 3 — Use `clinic.js` for Production-Like Analysis
`clinic heap` (part of the `clinic` package from NearForm) is the easiest way to get an annotated flame graph of heap growth under load.

```bash
npm install -g clinic
clinic heap -- node server.js
# apply load with autocannon or wrk
# ctrl+c, clinic opens a browser report
```

### Step 4 — Check for Listener Leaks
Node will warn you when an emitter has more than 10 listeners attached. Watch for:
```
MaxListenersExceededWarning: Possible EventEmitter memory leak detected
```

You can also audit at runtime:
```js
require('events').EventEmitter.defaultMaxListeners = 50; // raises the warning threshold
// or check specific emitters:
console.log(emitter.listenerCount('data'));
```

### Step 5 — Track Allocations with `--heap-prof`
Available since Node 12. Runs a sampling heap profiler and writes a `.heapprofile` file you can load in Chrome DevTools.

```bash
node --heap-prof server.js
```

This gives you allocation sites (the call stack where memory was allocated), which is often more useful than snapshots for pinpointing where leaky objects are created.

### Step 6 — Isolate with a Minimal Reproduction
Comment out subsystems one at a time (DB layer, cache, external HTTP calls, message queue consumers) and run the service under synthetic load. When the leak stops, the last thing you commented out is the source.

---

## Quick Checklist

- [ ] All event listeners removed when no longer needed (`removeListener` / `once`)
- [ ] All intervals and timeouts cleared when their owner is destroyed
- [ ] Caches have a max size and eviction policy
- [ ] Streams are always either consumed or destroyed, including on error paths
- [ ] No closures capturing large objects unintentionally
- [ ] Database connection pools are bounded and connections returned on error
- [ ] Promises have rejection handlers so they don't hang indefinitely

---

## Tools Summary

| Tool | Use Case |
|---|---|
| `node --inspect` + Chrome DevTools | Heap snapshots, comparison view |
| `node --heap-prof` | Allocation call-site profiles |
| `clinic heap` | Easy production-like heap analysis |
| `process.memoryUsage()` | Lightweight continuous monitoring |
| `heapdump` npm package | Programmatic snapshot triggers |
| `memwatch-next` | Leak event detection in code |
