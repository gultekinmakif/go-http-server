# Rulesets

GitHub branch-protection rulesets, version-controlled as JSON. GitHub does **not** auto-apply these from `.github/` — you import them manually (or scripted) per repo.

## Available rulesets

- [`pr-protection.json`](pr-protection.json) — protects the default branch: restrict deletions, block force pushes, require a PR with ≥1 approving review (code-owner review required, last-push approval required, stale reviews dismissed on push). Admin role bypasses (`actor_id: 5`).

## Importing a ruleset into a new repo

After creating a repo from this template:

```bash
# from the new repo's working tree
gh api --method POST \
  -H "Accept: application/vnd.github+json" \
  "repos/{owner}/{repo}/rulesets" \
  --input <(jq 'del(.id, .source, .source_type)' .github/rulesets/pr-protection.json)
```

The `jq` step strips the source-repo metadata so GitHub treats this as a fresh ruleset on the target repo. Requires the `gh` CLI authed with `repo` scope.

## Updating a ruleset

After tweaking the JSON in a checkout, re-import as an update:

```bash
RULESET_ID=$(gh api "repos/{owner}/{repo}/rulesets" --jq '.[] | select(.name=="PR Protection").id')
gh api --method PUT "repos/{owner}/{repo}/rulesets/$RULESET_ID" \
  --input <(jq 'del(.id, .source, .source_type)' .github/rulesets/pr-protection.json)
```

## Exporting current state

To capture the live ruleset back to JSON (e.g. after editing it in the UI):

```bash
RULESET_ID=$(gh api "repos/{owner}/{repo}/rulesets" --jq '.[] | select(.name=="PR Protection").id')
gh api "repos/{owner}/{repo}/rulesets/$RULESET_ID" > .github/rulesets/pr-protection.json
```
