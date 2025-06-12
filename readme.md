# changing

> A simple command-line tool to watch for changes in files or directories and exit the process upon detection.

## description

`changing` monitors specified paths for file modifications (e.g., saves,
deletions).Unlike continuous watchers, it exits with a success code (0)
as soon as a change is detected, making it useful for scripts that need
to react to changes.


## install

```sh
go install github.com/balazs4/changing
```

## usage

```sh
changing [-verbose] <FILE> [<FILE>]
# where <FILE> can be just `-` which means it reads from stdin
```

### examples

```sh
while changing readme.md; do fmt readme.md; done
```

```sh
while git ls-files | changing -; do go fmt .; done
```

## license

MIT

## author

balazs4
