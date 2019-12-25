# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## v1.1.1 - 2019-12-25
### Changed
- Switch back to using [Kingpin](https://github.com/alecthomas/kingpin) for
  command-line parsing.
- Make Bonclay compatible with the `go get` command.

### Removed
- Vendoring for dependencies.

## v1.1.0 - 2019-02-08
### Added
- `init` command which can be used to create a config file in the current
  directory.

### Changed
- Better error handling.

## v1.0.1 - 2018-11-15
### Changed
- `filepath.Walk()` is used for copying recursively instead of manual recursion.
  The functionality is the same as before but the code looks a bit nicer now.
- Command-line flags and arguments are parsed manually using the standard
  library instead of [Kingpin](https://github.com/alecthomas/kingpin).

## v1.0.0 - 2018-11-02

Initial release.
