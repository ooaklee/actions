<!-- action-docs-header source="action.yml" -->
# What The Ref
<!-- action-docs-header source="action.yml" -->

<!-- action-docs-description source="action.yml" -->
## Description

A GitHub Action that allows users to find the ref (sha/branch & path relative to workflow) of a named action that is utilised in a workflow
<!-- action-docs-description source="action.yml" -->

<!-- action-docs-runs source="action.yml" -->
## Runs

This action is a `node20` action.
<!-- action-docs-runs source="action.yml" -->

<!-- action-docs-inputs source="action.yml" -->
## Inputs

| name | description | required | default |
| --- | --- | --- | --- |
| `action-name` | <p>The name of the action to find the ref for</p> | `true` | `""` |
| `action-home-path-override` | <p>The override path to the action home directory</p> | `false` | `""` |
| `action-full-actions-store-path-override` | <p>The override path to the full path to which actioned are stored when referenced by in a workflow. By default it is set to <code>/home/runner/_work/_actions</code> (for the majority of non GitHub-hosted runners) or <code>/home/runner/work/_actions</code> (for GitHub-hosted runners i.e. ubuntu-latest). If this value is set, it will be used over action-home-path-override</p> | `false` | `""` |
<!-- action-docs-inputs source="action.yml" -->

<!-- action-docs-outputs source="action.yml" -->
## Outputs

| name | description |
| --- | --- |
| `ref` | <p>The identified reference (branch, commit or tag) of the specified action</p> |
| `path` | <p>The identified path of the specified action</p> |
<!-- action-docs-outputs source="action.yml" -->
