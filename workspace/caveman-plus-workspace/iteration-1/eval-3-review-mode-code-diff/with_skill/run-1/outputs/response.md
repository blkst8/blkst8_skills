L14: 🔵 nit: `{ id: id }` → `{ id }`. Shorthand is identical; the original was correct.
L15: 🟡 risk: `user` can be null if `findOne` returns nothing; `.profile.email` throws. Guard before access.
L18-21: 🔴 bug: `deleteUser` has no authorization check. Any caller can delete any user by id — this is `auth.js`.
L19: 🟡 risk: no error handling on `db.users.delete`; if it rejects, caller gets an unhandled rejection.
L20: 🔵 nit: `console.log` in production path. Remove or replace with a proper logger.
