# Bonclay

[![GitHub release](https://img.shields.io/github/release/talal/bonclay.svg)](https://github.com/talal/bonclay/releases/latest)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://travis-ci.org/talal/bonclay.svg?branch=master)](https://travis-ci.org/talal/bonclay)
[![Go Report Card](https://goreportcard.com/badge/github.com/talal/bonclay)](https://goreportcard.com/report/github.com/talal/bonclay)
[![GoDoc](https://godoc.org/github.com/talal/bonclay?status.svg)](https://godoc.org/github.com/talal/bonclay)

Bonclay is a fast and minimal backup tool.

Bonclay uses its configuration spec to backup, restore, or sync the specified files/directories.

## Installation

Download the latest pre-compiled binary from the [releases page](https://github.com/talal/bonclay/releases/latest).

Alternatively, you can also build from source:

The only required build dependency is [Go](https://golang.org/).

```
$ go get github.com/talal/bonclay
$ cd $GOPATH/src/github.com/talal/bonclay
$ make install
```

this will put the binary in `/usr/bin/bonclay` or `/usr/local/bin/bonclay` for macOS.

## Usage

Bonclay requires a configuration file in [YAML format](http://yaml.org). A minimal complete configuration could look like this:

```yaml
backup:
  overwrite: true

restore:
  overwrite: false

sync:
  clean: true
  overwrite: true

spec:
  ~/file: testfile
  ~/example/directory: test-dir
  ~/example/dir with space in name/file: test-dir2/testfile2
```

**Note**: Bonclay interprets the path for the sources and targets relative to its configuration file's location.

---

Bonclay has three operations:
- Sync
- Backup
- Restore

You can use different configuration files for different operations.

### Sync

Sync creates symbolic links between `source:target` pairs. It is useful for maintaining multiple copies of the same file without using up storage for those copies.

#### Options

| Option | Default value | Description |
| --- | --- | --- |
| clean | false | Remove broken symbolic links in the source's parent directory. |
| overwrite | false | Overwrite existing source file/directory. |

Default behavior:
- If the source's parent directory doesn't exist then create it.
- If there are any broken symbolic links in the source's parent directory then leave them be.
- If the source is already a symbolic link then remove that link and create a new link.
- If the source is an existing file or directory then skip it unless `overwrite: true`.

### Backup

Backup copies the sources to the targets. It is useful for backing up specific files/directories to a custom directory hierarchy.

#### Options

| Option | Default value | Description |
| --- | --- | --- |
| overwrite | false | Overwrite existing source. |

Default behavior:
- If the destination's parent directory doesn't exist then create it.
- If the destination is an existing file, directory, or symlink then skip it unless `overwrite: true`.

### Restore

Restore is the reverse of backup. It copies (read: restores) the targets to sources, and is useful for restoring specific files/directories to a custom directory hierarchy.

#### Options

| Option | Default value | Description |
| --- | --- | --- |
| overwrite | false | Overwrite existing source. |

Default behavior:
- If the destination's parent directory doesn't exist then create it.
- If the destination is an existing file, directory, or symlink then skip it unless `overwrite: true`.

## Example

Take a look at my [dotfiles repo](https://github.com/talal/dotfiles) and the included `bonclay.conf.yaml` file. I use Bonclay to keep my dotfiles in sync across different machines.

In order to use my dotfiles, you would only need to do the following:

```
$ cd /path/to/dotfiles
$ bonclay sync bonclay.conf.yaml
```

This will link the configuration/preference files specified in the configuration spec to the ones in the dotfiles directory.

**Note**: If you do intend to use my dotfiles then you might need to install some additional tools that my dotfiles depend on.
