# ops-agent

## Getting started

To build the Agent you need:
 * [Go](https://golang.org/doc/install) 1.19 or later. You'll also need to set your `$GOPATH` and have `$GOPATH/bin` in your path.
 * CMake version 3.12 or later and a C++ compiler

## Make

Parse dependency files.
```
make tidy
```

Generate amd64 *.gz package
```
make amd64

make linux_amd64.tar.gz
```

Generate x86_64 *.rpm package
```
make x86_64.rpm
```

## Run

You can run the agent with:
```
./opsAgent run
```
