# Changelog

## Unreleased

- Update Node.js 24.16.0 to 24.22.1 and Go 1.26.4 to 1.26.10; retain Yarn 1.22.22 and the Node 24 action runtime.
- Replace `action-docs` 2.5.1 and its dependency chain with a local generator using `yaml` 2.9.1 for action README tables. Update `lint-staged` 17.0.7 to 17.6.0, `prettier` 3.8.3 to 3.9.10, `go-githubactions` 1.3.2 to 1.4.0 and `testify` 1.11.1 to 1.12.1; replace transitive `gopkg.in/yaml.v3` 3.0.1 with `go.yaml.in/yaml/v3` 3.0.5, and rebuild the generated binaries.
- Remove the vulnerable `brace-expansion`, `minimatch`, `showdown` and `figlet` chains; update the retained `picomatch` 4.0.4 to 4.0.7 through `lint-staged`.

- Make the pre-commit hook portable to POSIX shells and add generator fixture checks for metadata escaping and empty defaults. Include README consistency checks in the normal test command.

## Earlier history

- Added a Go-backed GitHub Action template and fixed failure reporting so unsuccessful runs fail the workflow.
- Modernised the template toolchain and dependencies, introduced a Node 24 wrapper and checked-in Linux binaries, and documented the development workflow.

The repository has no recorded release tags for this history.
