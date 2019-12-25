# User guide

Bonclay requires a configuration file in [YAML format](http://yaml.org). A
minimal complete configuration looks like this:

```yaml
backup:
  overwrite: false

restore:
  overwrite: false

sync:
  clean: true
  overwrite: true

spec:
  ~/examplefile: file
  ../../examplefile: ../file
  ~/example dir: dir
  ../example dir/some other dir: ../new dir
```

The configuration file contains:
- Options for different Bonclay operations.
- `spec` section, which contains `source:target` pairs used by Bonclay to
  perform a specific operation. You can either use absolute paths or relative
  paths (relative to config file location) in the `spec` section (as shown
  above).

Since the configuration file is provided as an argument at the command line
therefore you can use different configuration files for different tasks.

For convenience, you can use the `init` command to create a configuration file
with sane defaults in the current directory:

```
$ bonclay init
```

## Operations

Bonclay has three operations:
- Backup
- Restore
- Sync

### Backup

Backup uses the 'source:target' pairs defined in the `spec` and copies the
sources to the targets. It is useful for backing up specific files/directories
to a custom directory hierarchy.

#### Options

| Option | Default value | Description |
|  --- | --- | --- |
| `overwrite` | false | Overwrite existing file/directory. |

Default behavior:
- If the destination's parent directory doesn't exist then create it.
- If the destination is an existing file, directory, or symlink then skip it
  unless `overwrite: true`.

### Restore

Restore is the reverse of backup. It copies the targets to sources and is
useful for restoring specific files/directories to a custom directory hierarchy
(or back to their original location, if same `spec` was used for backup).

#### Options

| Option | Default value | Description |
| --- | --- | --- |
| `overwrite` | false | Overwrite existing file/directory. |

Default behavior:
- If the destination's parent directory doesn't exist then create it.
- If the destination is an existing file, directory, or symlink then skip it
  unless `overwrite: true`.

### Sync

Sync creates symbolic links between `source:target` pairs. It is useful for
maintaining multiple copies of the same file/directories without using up
storage for those copies.

#### Options

| Option | Default value | Description |
| --- | --- | --- |
| `clean` | false | Remove broken symbolic links in the source's parent directory. |
| `overwrite` | false | Overwrite existing file/directory. |

Default behavior:
- If the source's parent directory doesn't exist then create it.
- If there are any broken symbolic links in the source's parent directory then
  leave them be, unless `clean: true`.
- If the source is already a symbolic link then remove the existing link and
  create a new link.
- If the source is an existing file or directory then skip it unless
  `overwrite: true`.
