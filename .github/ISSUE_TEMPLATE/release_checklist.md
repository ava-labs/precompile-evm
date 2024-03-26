---
name: Release Checklist
about: Create a ticket to track a release
title: ""
labels: release
assignees: ""
---

## Release

The release version and a description of the planned changes to be included in the release.

## Issues

Link the major issues planned to be included in the release.

## Documentation

Link the relevant documentation PRs for this release.

## Checklist

- [ ] Update Precompile-EVM version in plugin/main.go
- [ ] Bump AvalancheGo dependency for RPCChainVM Compatibility in go.mod and versions.sh
- [ ] Bump Subnet-EVM version in go.mod and versions.sh
- [ ] Add new entry in compatibility.json for RPCChainVM Compatibility
- [ ] Update AvalancheGo compatibility in README
- [ ] Update hello-world-example branch with latest master, run CI tests
