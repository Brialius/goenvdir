# goenvdir
[![Go Report Card](https://goreportcard.com/badge/github.com/Brialius/goenvdir)](https://goreportcard.com/report/github.com/Brialius/goenvdir)
## Usage
```
Usage:
goenvdir dir child
        dir
                directory with files named as env variables
        child
                executable program with all parameters
```
## Build
### make goals
|Goal|Description|
|----|-----------|
|setup|download and install required dependencies|
|test|run tests|
|build|build binary: `bin/goenvdir` or `bin/goenvdir.exe` for windows|
|install|install binary to `$GOPATH/bin`|
|lint|run linters|
|clean|run `go clean`|
|mod-refresh|run `go mod tidy` and `go mod vendor`|
|ci|run all steps needed for CI|
|version|show current git tag if any matched to `v*` exists|
|release|set git tag and push to repo `make release ver=v1.2.3`|