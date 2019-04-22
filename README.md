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

If you want to compile the code to a C shared library, you have 2 ways:

### Automatic (recommended)

Git commit with a message containing the string "[compile]". If you have nothing to commit but still want to trigger the compilation process, you can create an empty commit:

```
git commit --allow-empty -m '[compile]'
```

Go to Travis CI website to [see the status of your compilation process](https://travis-ci.org/asyncapi/parser/builds).

If there are changes in the binaries, Travis will push them to your branch. Beware it will push a commit for each platform (Windows, Mac, Linux).

### Manual

Compiling to C is rather trivial but you have to make sure you install everything you'll need. Check out the requirements for each platform:

#### Windows

Install Windows Build Tools (Node.js required):

```
npm install --global windows-build-tools
```

This should install Visual C++ Build Tools and Python.

And then run:

```
compile-windows
```

#### Mac OS X

Install Xcode Command Line Tools:

```
xcode-select --install
```

And then run:

```
./compile-darwin.sh
```

#### Linux

Install GCC:

```
sudo apt install gcc
```

```
sudo apt install build-essential
```

And then run:

```
./compile-linux.sh
```

## Authors

* Fran Mendez
* Raisel Melian
* Ruben Hervas
