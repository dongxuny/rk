# rk
RK command line tools. It contains couple of useful utility functionalities including downloading dependencies, build go project etc.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Command Overview](#command-overview)
  - [rk help](#rk-help)
  - [rk install](#rk-install)
  - [rk uninstall](#rk-uninstall)
  - [rk ut](#rk-ut)
  - [rk build](#rk-build)
  - [rk clear](#rk-clear)
  - [rk pack](#rk-pack)
  - [rk docker](#rk-docker)
  - [rk run](#rk-run)
- [build.yaml](#buildyaml)
  - [Example](#example)
- [Development Status: Stable](#development-status-stable)
- [Contributing](#contributing)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Installation
```shell script
go get -u github.com/rookie-ninja/rk/cmd
```

## Quick Start
We'll start with a general overview of the commands. 
There are more commands, and we will get into] usage below, but this shows the basic functionality.

```shell script
COMMANDS:
   build      Build project which contains build.yaml
   clear      Clear target folder generated by rk build
   docker     Build a docker image built with rk build command
   install    Install third-party software
   uninstall  Uninstall third-party software
   pack       Pack target folder generated by rk build
   run        Run server build by rk build
   ut         Run unit test
   help, h    Shows a list of commands or help for one command
```

## Command Overview

### rk help
```shell script
rk help
```
Print help message

### rk install
```shell script
rk install 
```

Subcommands
```shell script
COMMANDS:
   buf                      install buf on local machine
   cfssl                    install cfssl on local machine
   cfssljson                install cfssljson on local machine
   gocov                    install gocov on local machine
   golangci-lint            install golangci-lint on local machine
   mockgen                  install mockgen on local machine
   pkger                    install pkger on local machine
   protobuf                 install protobuf on local machine
   protoc-gen-doc           install protoc-gen-doc on local machine
   protoc-gen-go            install protoc-gen-go on local machine
   protoc-gen-grpc-gateway  install protoc-gen-grpc-gateway on local machine
   protoc-gen-openapiv2     install protoc-gen-openapiv2 on local machine
   swag                     install swag on local machine
   rk-std                   install rk standard environment on local machine
   help, h                  Shows a list of commands or help for one command
```

### rk uninstall
```shell script
rk uninstall 
```

Subcommands
```shell script
COMMANDS:
   buf                      uninstall buf from local machine
   cfssl                    uninstall cfssl from local machine
   cfssljson                uninstall cfssljson from local machine
   gocov                    uninstall gocov from local machine
   golangci-lint            uninstall golangci-lint from local machine
   mockgen                  uninstall mockgen from local machine
   pkger                    uninstall pkger from local machine
   protobuf                 uninstall protobuf from local machine
   protoc-gen-doc           uninstall protoc-gen-doc from local machine
   protoc-gen-go            uninstall protoc-gen-go from local machine
   protoc-gen-grpc-gateway  uninstall protoc-gen-grpc-gateway from local machine
   protoc-gen-openapiv2     uninstall protoc-gen-openapiv2 from local machine
   swag                     uninstall swag from local machine
   rk-std                   uninstall rk standard environment on local machine
   help, h                  Shows a list of commands or help for one command
```

### rk ut
```shell script
rk ut # Run unit test
```

Subcommands
```shell script
COMMANDS:
   help, h  Shows a list of commands or help for one command
```

### rk build
```shell script
rk build - Build project which contains build.yaml
```

### rk clear
```shell script
rk clear - Clear target folder generated by rk build
```

### rk pack
```shell script
rk pack - Pack target folder generated by rk build
```

### rk docker
```shell script
rk docker - Build a docker image built with rk build command
```

Subcommands
```shell script
COMMANDS:
   build    build a docker image built with rk build command
   run      run a docker image built with rk docker build
   help, h  Shows a list of commands or help for one command
```

### rk run
```shell script
rk run - run server build by rk build
```

## build.yaml
Bellow commands need to build project as RK style first in order to continue.
- rk build
- rk run
- rk pack
- rk docker build
- rk docker run

### Example
```yaml
---
build:
  type: go                          # Optional, default:go
  main: "internal/main.go"          # Optional, default: ./main.go
  GOOS: ""                          # Optional, default: current OS
  GOARCH: ""                        # Optional, default: current Arch
  args: ""                          # Optional, default: "", arguments which will attached to [go build] command
  copy: [""]                        # Optional, default: [], directories or files need to copy to [target] folder
  commands:                         
    before: []                      # Optional, default: [], commands would be invoked before [go build] command locally
    after: []                       # Optional, default: [], commands would be invoked after [go build] command locally
  scripts:
    before: []                      # Optional, default: [], scripts would be executed before [go build] command locally
    after: []                       # Optional, default: [], scripts would be executed after [go build] command locally
docker:
  build:
    registry: ""                    # Optional, default: [package name]
    tag: ""                         # Optional, default: [current git tag or branch-latestCommit]
    args: [""]                      # Optional, default: "", docker args which will be attached to [docker build] command
  run:
    args: ["-p", "8080:8080"]       # Optional, default: "", docker args which will be attached to [docker run] command
```

## Development Status: Stable

## Contributing
We encourage and support an active, healthy community of contributors &mdash;
including you! Details are in the [contribution guide](CONTRIBUTING.md) and
the [code of conduct](CODE_OF_CONDUCT.md). The rk maintainers keep an eye on
issues and pull requests, but you can also report any negative conduct to
dongxuny@gmail.com. That email list is a private, safe space; even the zap
maintainers don't have access, so don't hesitate to hold us to a high
standard.

Released under the [MIT License](LICENSE).
