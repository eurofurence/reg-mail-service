# reg-mail-service

<img src="https://github.com/eurofurence/reg-mail-service/actions/workflows/go.yml/badge.svg" alt="test status"/>
<img src="https://github.com/eurofurence/reg-mail-service/actions/workflows/codeql-analysis.yml/badge.svg" alt="code quality status"/>

## Overview

A backend service...

Implemented in go.

Command line arguments
```-config <path-to-config-file> [-migrate-database]```

## Installation

This service uses go modules to provide dependency management, see `go.mod`.

If you place this repository OUTSIDE of your gopath, `go build main.go` and `go test ./...` will download all
required dependencies by default.
