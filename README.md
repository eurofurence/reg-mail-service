# reg-mail-service

<img src="https://github.com/eurofurence/reg-mail-service/actions/workflows/go.yml/badge.svg" alt="test status"/>

## Overview

Mail backend service.

Implemented in go.

Command line arguments
```-config <path-to-config-file> [-migrate-database]```

## Installation

This service uses go modules to provide dependency management, see `go.mod`.

Copy config.example.yaml from the docs/ directory into the project root. Rename that file into:
config.yaml
Then configure all variables as needed.

If you place this repository OUTSIDE of your gopath, `go build main.go` and `go test ./...` will download all
required dependencies by default.

## Open Issues and Ideas

We track open issues as GitHub issues on this repository once it becomes clear what exactly needs to be done.
