## Code Review: `auth.js`

### `getUser` — Minor Style Issue

```diff
-  const user = await db.users.findOne({ id });
+  const user = await db.users.findOne({ id: id });
```

This change is functionally equivalent, but it's a **regression in style**. The shorthand `{ id }` is idiomatic modern JavaScript (ES6+) and should be preferred over the verbose `{ id: id }`. This change should be reverted.

Additionally, `user.profile.email` is accessed without a null check. If `findOne` returns `null` (user not found), this will throw a `TypeError`. Consider guarding:

```js
if (!user) throw new Error(`User ${id} not found`);
```

---

### `deleteUser` — Several Issues

```js
async function deleteUser(id) {
  await db.users.delete({ id });
  console.log('user deleted');
}
```

**1. `console.log` left in production code**
`console.log('user deleted')` is debug output. This should be removed or replaced with structured/audit logging appropriate for a production auth module.

**2. No authorization check**
This is an auth module. Deleting a user is a privileged, destructive operation. There's no check that the caller has permission to delete the user. Missing an `authorizeRequest` call or similar guard here is a significant security concern.

**3. No error handling**
If `db.users.delete` fails, the error will propagate uncaught to the caller with no context. At minimum, consider wrapping in a try/catch and re-throwing with a descriptive message, or let the caller handle it — but document that behavior.

**4. No verification the user exists**
Deleting a non-existent user likely silently succeeds. Whether that's acceptable depends on your DB layer, but it may mask bugs (e.g., wrong ID passed in).

**5. No return value**
Not strictly required, but returning a success indicator or the deleted record count can help callers confirm the operation had an effect.

---

### Summary

| Issue | Severity |
|---|---|
| `{ id: id }` instead of `{ id }` | Low |
| Missing null check on `user.profile.email` | Medium |
| `console.log` in production auth code | Medium |
| No authorization check in `deleteUser` | High |
| No error handling in `deleteUser` | Medium |
| No existence check before delete | Low |
