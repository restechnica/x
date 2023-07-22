# X

[![github.com release badge](https://img.shields.io/github/release/restechnica/x.svg)](https://github.com/restechnica/x/)
[![github.com workflow badge](https://github.com/restechnica/x/workflows/main/badge.svg)](https://github.com/restechnica/x/actions?query=workflow%3Amain)
[![go.pkg.dev badge](https://pkg.go.dev/badge/github.com/restechnica/x)](https://pkg.go.dev/github.com/restechnica/x)
[![goreportcard.com badge](https://goreportcard.com/badge/github.com/restechnica/x)](https://goreportcard.com/report/github.com/restechnica/x)
[![img.shields.io MPL2 license badge](https://img.shields.io/github/license/restechnica/x)](./LICENSE)

A modern day task runner which tries to be language independent.

## Table of Contents

* [Requirements](#requirements)
* [How to install](#how-to-install)
    * [Github](#github)
    * [Homebrew](#homebrew)
    * [Golang](#golang)
* [Usage](#usage)

## Requirements

Go 1.20

## How to install

`x` can be retrieved from GitHub or a Homebrew tap. Run `x -h` to validate the installation.
The tool is available for Windows, Linux and macOS.

### github

`x` is available through github. The following example works for a GitHub Workflow, other CI/CD tooling will require a different path setup.

```shell
SEMVERBOT_VERSION=1.0.0
mkdir bin
echo "$(pwd)/bin" >> $GITHUB_PATH
curl -o bin/x -L https://github.com/restechnica/semverbot/releases/download/v$SEMVERBOT_VERSION/sbot-linux-amd64
chmod +x bin/x
```

### homebrew

`x` is available through the public tap [github.com/restechnica/homebrew-tap](https://github.com/restechnica/homebrew-tap)

```shell
brew tap restechnica/tap git@github.com:restechnica/homebrew-tap.git
brew install restechnica/tap/x
```

### golang

`x` is written in golang, which means you can use `go install`. Make sure the installation folder, which depends on your golang setup, is in your system PATH.

```shell
go install github.com/restechnica/semverbot/cmd/x@v1.0.0
```

## Usage

Each command has a `-h, --help` flag available. Support for `-v, --verbose` and `-d, --debug` has been added as well.

More docs coming soon.