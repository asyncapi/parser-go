# AsyncAPI Parser

# THIS IS A WORK IN PROGRESS

It parses AsyncAPI documents.

## Trying it from the command line

In order to test the parser for the command line you can run:

```go
go run cmd/cli/main.go -file <PATH_TO_DEFINITION_FILE>.yaml
```

This will output the parsed definition in json format or error in case of invalid definition file.

## Compiling

We use [xgo](https://github.com/karalabe/xgo) to compile for multiple platforms.

### Installing xgo

```bash
docker pull karalabe/xgo-latest
go get github.com/karalabe/xgo
```

If you want to know more about the installation process, check out [xgo documentation](https://github.com/karalabe/xgo#installation).

### Running compilation script

```bash
./compile.sh
```

Resulting binaries and `.h` files will be placed in the `bin` directory of the project.

## Authors

* Fran Mendez
* Raisel Melian
* Ruben Hervas
