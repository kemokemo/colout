# colout (colored output tool)

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![codecov](https://codecov.io/gh/kemokemo/colout/branch/master/graph/badge.svg)](https://codecov.io/gh/kemokemo/colout) [![Go Report Card](https://goreportcard.com/badge/github.com/kemokemo/colout)](https://goreportcard.com/report/github.com/kemokemo/colout)

A small and simple CLI tool that makes it easy to use colored output in scripts and other applications.

## Usage

```sh
$ colout -h
A CLI tool that makes it easy to use colored output.

Usage:
  colout [command]

Available Commands:
  error       colored red output
  help        Help about any command
  info        colored blue output
  success     colored green output
  version     Print the version number
  warn        colored yellow output

Flags:
  -h, --help   help for colout

Use "colout [command] --help" for more information about a command.
```

## Examples

### Both args and pipe are available

The following two samples both have the same output.

```sh
colout info "Hey! What's up?"
```

```sh
echo "Hey! What's up?" | colout info
```

### Shell script

<details>

```sh
#!/bin/bash

dosomething

if [ $? = 0 ]; then
  colout success "successfully done!"
else
  colout error "failed to run 'dosomething'."
fi
```
</details>

### Power shell

<details>

```sh
dosomething

if ( $? ){
  colout success "successfully done!"
}
else {
  colout error "failed to run 'dosomething'."
}
```
</details>

### Windows batch

<details>

```sh
@echo off

dosomething

if %errorlevel% == 0 (
  colout success "successfully done!"
) else (
  colout error "failed to run 'dosomething'."
) 
```
</details>

## Install

### Binary

Get the latest version from [the release page](https://github.com/kemokemo/colout/releases/latest), and download the archive file for your operating system/architecture. Unpack the archive, and put the binary somewhere in your $PATH.

### Build yourself

```sh
go get github.com/kemokemo/colout
```

## License

[MIT](https://github.com/kemokemo/colout/blob/master/LICENSE)

## Author

[kemokemo](https://github.com/kemokemo)

