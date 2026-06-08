---
id: adrs-adr001
title: 'ADR001: Modernise Go Action Toolchain'
# prettier-ignore
description: Architecture Decision Record (ADR) for modernising the Go action template toolchain, dependencies, runtime, and generated binaries
---

## Context

This repository provides reusable GitHub Action templates. The Go action template is implemented as a JavaScript action that invokes prebuilt Go binaries for Linux AMD64 and Linux ARM64 runners.

The template previously used older language and package versions. Node.js 20 is no longer the safest long-term runtime choice for this template, while the GitHub Actions metadata syntax reference lists `node24` as a supported JavaScript action runtime. The repository also had more than one `.tool-versions` file, which allowed nested toolchain pins to drift from the repository-level development environment. Some JavaScript transitive dependencies in the old lockfile had known advisories, and the Go module directive used an invalid patch-level `go` version format for modern Go tooling.

The template needs to stay easy for action authors to copy, test, document, and publish. Any source-level change to the Go action must also refresh the checked-in `dist/` binaries because workflow consumers execute those binaries through `invoke-binary.js`.

## Decision

We will use the repository root `.tool-versions` file as the single source of truth for local toolchains.

We will pin the template to current safe language versions and package versions:

- Node.js 24.16.0 for local JavaScript tooling and `runs.using: node24` for the action runtime.
- Yarn 1.22.22 for the existing Yarn classic workspace.
- Go 1.26.4 as the toolchain, with `go 1.26` in `go.mod`.
- Current safe JavaScript and Go dependencies as recorded in `package.json`, `yarn.lock`, `go.mod`, and `go.sum`.

We will keep the Go action template structure:

- `action.yml` defines the public action API.
- `invoke-binary.js` remains the JavaScript action entrypoint.
- `src/internal/config` parses and validates action inputs.
- `src/internal/runner` contains the action behaviour.
- `dist/action-amd64` and `dist/action-arm64` are regenerated from Go source with `yarn package`.

We will add explicit documentation comments and README guidance so authors can understand the template boundaries and adapt the action safely.

## Consequences

The template now uses supported language runtimes and avoids known vulnerable JavaScript transitive dependencies from the previous lockfile.

Developers can install and use one repository-level asdf toolchain instead of reconciling nested `.tool-versions` files.

Action consumers will run the refreshed Linux binaries through the `node24` wrapper.

The template is clearer for future authors because input parsing, runner behaviour, binary selection, and adaptation steps are documented.

The repository remains on Yarn classic. This avoids a package-manager migration in the same change, but it means future maintainers should continue to account for Yarn classic behaviour and warnings when running on newer Node.js versions.

Checked-in binaries will continue to change whenever Go source, Go compiler versions, or build flags change. Maintainers must keep source and `dist/` artefacts synchronised before publishing.
