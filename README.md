# Archived project. No maintenance.

**What does this mean?** No further code will be committed and the repository
will remain in a read-only state. [Releases][releases] will still be available
for download. Any existing projects that use Bonclay will continue to work.

**Current status:** Bonclay is stable (as far as I know).

**Why?** I no longer use Bonclay personally, I've switched to using
[`chezmoi`](https://github.com/twpayne/chezmoi) for dotfiles and
[Vorta](https://github.com/borgbase/vorta) for backups.

# Bonclay

[![GitHub Release](https://img.shields.io/github/release/talal/bonclay.svg?style=flat-square)](https://github.com/talal/bonclay/releases/latest)
[![Build Status](https://img.shields.io/travis/talal/bonclay/master.svg?style=flat-square)](https://travis-ci.org/talal/bonclay)
[![Go Report Card](https://goreportcard.com/badge/github.com/talal/bonclay?style=flat-square)](https://goreportcard.com/report/github.com/talal/bonclay)
[![Software License](https://img.shields.io/github/license/talal/bonclay.svg?style=flat-square)](LICENSE)

Bonclay is a fast and minimal backup tool.

Bonclay uses a yaml file that has `source:target` pairs to backup, restore, or sync the specified files/directories.

The following is a demo on how you can use Bonclay to manage your dotfiles:

[![asciicast](https://asciinema.org/a/226247.svg)](https://asciinema.org/a/226247)

Refer to use [user guide](./docs/guide.md) for instructions.

## Installation

### Installer script

The simplest way to install Bonclay on Linux or macOS is to run:

```
$ sh -c "$(curl -sL git.io/getbonclay)"
```

This will put the binary in `/usr/local/bin/bonclay`

### Pre-compiled binaries

Pre-compiled binaries for Linux and macOS are avaiable on the
[releases page][releases].

[releases]: https://github.com/talal/bonclay/releases/latest

The binaries are static executables.

### Homebrew

```
$ brew install talal/tap/bonclay
```

### Building from source

The only required build dependency is [Go](https://golang.org/).

```
$ go get github.com/talal/bonclay
$ cd $GOPATH/src/github.com/talal/bonclay
$ make install
```

This will put the binary in `/usr/local/bin/bonclay`

## Usage

Refer to the [user guide](./docs/guide.md).
