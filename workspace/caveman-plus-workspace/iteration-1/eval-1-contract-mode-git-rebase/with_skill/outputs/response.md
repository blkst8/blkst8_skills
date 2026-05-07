`git rebase -i HEAD~3` starts an interactive rebase on your last 3 commits. Your editor'll open with a list like this:

```
pick a1b2c3 Add login form
pick d4e5f6 Fix typo
pick 7g8h9i Finalize login styles
```

Each line is a commit, and you can change the action word to control what happens to it:

- **`pick`** — keep it as-is
- **`reword`** — keep the commit but edit its message
- **`squash`** — merge it into the commit above it (keeps the message)
- **`fixup`** — same as squash, but throws away its message
- **`edit`** — pause so you can amend the commit
- **`drop`** — delete it entirely

Save the file, close the editor, and Git'll replay the commits according to your instructions.

**Why you'd use it:**

You're probably using it to clean up a messy local history before merging or pushing. If you've been committing things like "wip", "fix", "actually fix", "ok now it works", you can squash those down into one meaningful commit. It's also useful for rewording a bad commit message, reordering commits, or splitting one big commit into smaller ones.

The key thing to know: it rewrites history. That's fine for commits that haven't been shared yet, but if you've already pushed these commits to a branch others're working on, you'll cause problems. Don't use it on shared history unless you've coordinated with your team.
