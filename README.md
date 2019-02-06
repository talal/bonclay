# Bonclay

[![GitHub Release](https://img.shields.io/github/release/talal/bonclay.svg?style=flat-square)](https://github.com/talal/bonclay/releases/latest)
[![Build Status](https://img.shields.io/travis/talal/bonclay/master.svg?style=flat-square)](https://travis-ci.org/talal/bonclay)
[![Go Report Card](https://goreportcard.com/badge/github.com/talal/bonclay?style=flat-square)](https://goreportcard.com/report/github.com/talal/bonclay)
[![Software License](https://img.shields.io/github/license/talal/bonclay.svg?style=flat-square)](LICENSE)


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
