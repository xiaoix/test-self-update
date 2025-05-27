# test-self-update

[![Build Status](https://github.com/xiaoix/test-self-update/workflows/test-unit/badge.svg)](https://github.com/xiaoix/test-self-update/actions?query=branch%3Amaster+workflow%3Atest-unit)
[![Coverage Status](https://codecov.io/gh/xiaoix/test-self-update/branch/master/graph/badge.svg)](https://codecov.io/gh/xiaoix/test-self-update)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/github.com/xiaoix/test-self-update)
[![Time Tracker](https://wakatime.com/badge/github/xiaoix/test-self-update.svg)](https://wakatime.com/badge/github/xiaoix/test-self-update)
![Code lines](https://sloc.xyz/github/xiaoix/test-self-update/?category=code)
![Comments](https://sloc.xyz/github/xiaoix/test-self-update/?category=comments)

<!--- TODO Update README.md -->

Project template with GitHub actions for Go.

## Install

```
go install github.com/xiaoix/test-self-update@latest
$(go env GOPATH)/bin/test-self-update --help
```

Or download binary from [releases](https://github.com/xiaoix/test-self-update/releases).

### Linux AMD64

```
wget https://github.com/xiaoix/test-self-update/releases/latest/download/linux_amd64.tar.gz && tar xf linux_amd64.tar.gz && rm linux_amd64.tar.gz
./test-self-update -version
```

### Macos Intel

```
wget https://github.com/xiaoix/test-self-update/releases/latest/download/darwin_amd64.tar.gz && tar xf darwin_amd64.tar.gz && rm darwin_amd64.tar.gz
codesign -s - ./test-self-update
./test-self-update -version
```

### Macos Apple Silicon (M1, etc...)

```
wget https://github.com/xiaoix/test-self-update/releases/latest/download/darwin_arm64.tar.gz && tar xf darwin_arm64.tar.gz && rm darwin_arm64.tar.gz
codesign -s - ./test-self-update
./test-self-update -version
```


## Usage

Create a new repository from this template, check out it and run `./run_me.sh` to replace template name with name of
your repository.
