# Bonclay

[![GitHub release](https://img.shields.io/github/release/talal/bonclay.svg)](https://github.com/talal/bonclay/releases/latest)
[![Build Status](https://travis-ci.org/talal/bonclay.svg?branch=master)](https://travis-ci.org/talal/bonclay) 
[![Go Report Card](https://goreportcard.com/badge/github.com/talal/bonclay)](https://goreportcard.com/report/github.com/talal/bonclay)
[![GoDoc](https://godoc.org/github.com/talal/bonclay?status.svg)](https://godoc.org/github.com/talal/bonclay)

Bonclay is a simple dotfiles manager. Well... technically, it is a backup/restore/sync tool magiggy.

Bonclay uses its configuration spec to backup, restore, or sync the specified files/directories.

I call it a *dotfiles manager* because that is the use case I had, when I created it. But in reality, it is capable of much more :)

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

Bonclay has three operations:
- Sync
- Backup
- Restore

### Sync

Sync creates symbolic links between `source:target` pairs. It is useful for maintaining multiple copies of the same file without using up storage for those copies.

For a single `source:target` pair, the sync operation is equivalent to manually doing:

```
ln -s source target
```

#### Options

| Option | Default value | Description |
| --- | --- | --- |
| clean | false | Remove broken symbolic links in the source's parent directory. |
| overwrite | false | Overwrite existing source file/directory. |

Default behavior: 
- If there are any broken symbolic links in the source's parent directory then leave them be.
- If the source is already a symbolic link then remove that link and create a new link.
- If the source is an existing file or directory then skip it unless `overwrite: true`.

### Backup

Backup copies the sources to the targets. It is useful for backing up specific files/directories in a custom hierarchical order.

For a single `source:target` pair, the backup operation is equivalent to manually doing:

```
cp -r source target
```

#### Options

| Option | Default value | Description |
| --- | --- | --- |
| overwrite | false | Overwrite existing source. |

Default behavior: 
- If the target is an existing file, directory, or symlink then skip it unless `overwrite: true`.

### Restore

Restore is the reverse of backup. It copies the targets to sources, and is useful for restoring specific files/directories to a custom hierarchical order.

For a single `source:target` pair, the restore operation is equivalent to manually doing:

```
cp -r target source
```

#### Options

| Option | Default value | Description |
| --- | --- | --- |
| overwrite | false | Overwrite existing source. |

Default behavior: 
- If the source is an existing file, directory, or symlink then skip it unless `overwrite: true`.

## Example

Take a look at my [dotfiles repo](https://github.com/talal/dotfiles) and the included `bonclay.conf.yaml` file. I use Bonclay to keep my dotfiles in sync across different machines.
