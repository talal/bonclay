# User guide

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

The configuration file contains:
- Options for different Bonclay operations.
- `spec` section, which contains `source:target` pairs which are used by Bonclay to perform a specific operation. You can either use absolute paths or relative paths in the `spec` section (as shown above).

You can use multiple configuration files for different tasks.

## Operations

Bonclay has three operations:
- Sync
- Backup
- Restore

### Sync

Sync creates symbolic links between `source:target` pairs. It is useful for maintaining multiple copies of the same file without using up storage for those copies.

#### Options

| Option | Default value | Description |
| --- | --- | --- |
| `clean` | false | Remove broken symbolic links in the source's parent directory. |
| `overwrite` | false | Overwrite existing source file/directory. |

Default behavior:
- If the source's parent directory doesn't exist then create it.
- If there are any broken symbolic links in the source's parent directory then leave them be, unless `clean: true`.
- If the source is already a symbolic link then remove the existing link and create a new link.
- If the source is an existing file or directory then skip it, unless `overwrite: true`.

### Backup

Backup uses the 'source:target' pairs defined in the `spec` and copies the sources to the targets. It is useful for backing up specific files/directories to a custom directory hierarchy.

#### Options

| Option | Default value | Description |
| --- | --- | --- |
| `overwrite` | false | Overwrite existing source file/directory. |

Default behavior:
- If the destination's parent directory doesn't exist then create it.
- If the destination is an existing file, directory, or symlink then skip it, unless `overwrite: true`.

### Restore

Restore is the reverse of backup. It copies the targets to sources and is useful for restoring specific files/directories to a custom directory hierarchy.

#### Options

| Option | Default value | Description |
| --- | --- | --- |
| `overwrite` | false | Overwrite existing source file/directory. |

Default behavior:
- If the destination's parent directory doesn't exist then create it.
- If the destination is an existing file, directory, or symlink then skip it, unless `overwrite: true`.

## Example

Take a look at my [dotfiles repo](https://github.com/talal/dotfiles) and the included `bonclay.conf.yaml` file. I use Bonclay to keep my dotfiles in sync across different machines.

In order to use my dotfiles, you would only need to do the following:

```
$ cd /path/to/dotfiles
$ bonclay sync bonclay.conf.yaml
```

This will link the configuration/preference files specified in the configuration spec to the ones in the dotfiles directory.

**Note**: If you do intend to use my dotfiles then you will need to install some additional tools that my dotfiles depend on: `./homebrew/install`
