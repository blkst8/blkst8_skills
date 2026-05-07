# git rebase -i HEAD~3

Running `git rebase -i HEAD~3` opens an interactive rebase session that lets you edit the last 3 commits in your branch's history.

## What happens

Git opens your configured editor with a list of the 3 most recent commits, each prefixed with a command word — by default, `pick`. It looks something like this:

```
pick a1b2c3d Add login form
pick e4f5g6h Fix validation bug
pick i7j8k9l Update styles

# Rebase 0f1e2d3..i7j8k9l onto 0f1e2d3 (3 commands)
#
# Commands:
# p, pick = use commit
# r, reword = use commit, but edit the commit message
# e, edit = use commit, but stop for amending
# s, squash = meld into previous commit
# f, fixup = like squash, but discard this commit's log message
# d, drop = remove commit
```

You change the command words and save the file. Git then replays the commits according to your instructions.

## Why someone would use it

- **Squash commits together** — combine a series of "WIP" or fix-up commits into one clean commit before merging a pull request
- **Reword commit messages** — fix a typo or improve clarity in a past commit message
- **Reorder commits** — rearrange the order of commits (by reordering lines in the editor)
- **Drop commits** — remove a commit entirely (change `pick` to `drop`)
- **Split or amend a commit** — pause mid-rebase with `edit` to amend a commit or split it into multiple commits

The core motivation is keeping a clean, readable commit history — especially before pushing to a shared branch or opening a pull request. Instead of a history cluttered with "oops" and "fix typo" commits, you present a coherent narrative of changes.

## Important caveat

This rewrites commit history. Don't use it on commits that have already been pushed to a shared branch, as it forces others to reconcile diverged history.
