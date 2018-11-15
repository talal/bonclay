# v1.0.1 (2018-11-15)

Changes:
* filepath.Walk() is used for copying recursively instead of manual recursion. The functionality is the
	same as before but the code looks a bit nicer now.
* Command-line flags and arguments are parsed manually using the standard library instead of [kingpin](https://github.com/alecthomas/kingpin).

# v1.0.0 (2018-11-02)

Initial release.
