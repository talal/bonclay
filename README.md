# Bonclay

[![GitHub release](https://img.shields.io/github/release/talal/bonclay.svg)](https://github.com/talal/bonclay/releases/latest)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://travis-ci.org/talal/bonclay.svg?branch=master)](https://travis-ci.org/talal/bonclay)
[![Go Report Card](https://goreportcard.com/badge/github.com/talal/bonclay)](https://goreportcard.com/report/github.com/talal/bonclay)
[![GoDoc](https://godoc.org/github.com/talal/bonclay?status.svg)](https://godoc.org/github.com/talal/bonclay)

Bonclay is a fast and minimal backup tool.

Bonclay uses its configuration spec to backup, restore, or sync the specified files/directories.

## Installation

### Pre-compiled binaries

Pre-compiled binaries for Linux and macOS are avaiable on the [releases page](https://github.com/talal/bonclay/releases/latest).

The binaries are static executables.

### Homebrew

```
brew install talal/tap/bonclay
```

### Building from source

The only required build dependency is [Go](https://golang.org/).

```
$ go get github.com/talal/bonclay
$ cd $GOPATH/src/github.com/talal/bonclay
$ make install
```

this will put the binary in `/usr/bin/bonclay` or `/usr/local/bin/bonclay` for macOS.

## Usage

Take a look at the [user guide](./doc/guide.md).
