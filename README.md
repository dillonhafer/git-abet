# git-abet

Show frequently commited files with the current file

```
Usage:
  git-abet [OPTIONS] file

Flags:
  -h, --help         help for git-abet
  -n, --number int   Number of files to return (default 5)
```

Example:

```
$ git-abet -n 5 README.md
.gitignore
cmd/root.go
main.go
```
