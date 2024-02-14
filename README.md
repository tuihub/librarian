# Librarian

[![](https://img.shields.io/github/v/release/tuihub/librarian.svg)](https://github.com/tuihub/librarian/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/tuihub/librarian)](https://goreportcard.com/report/github.com/tuihub/librarian)
[![golangci-lint](https://github.com/TuiHub/Librarian/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/TuiHub/Librarian/actions/workflows/golangci-lint.yml)
[![codecov](https://codecov.io/gh/TuiHub/Librarian/branch/master/graph/badge.svg?token=9E9VIBWYZV)](https://codecov.io/gh/TuiHub/Librarian)

Librarian is the standard server implementation of TuiHub.  

It is recommended to use with the standard client implementation [Waiter](https://github.com/tuihub/waiter). Check [document](https://docs.tuihub.org) site for user guide.

## Usage

### Server binary

Build the server with `make`,
or download the binary file in [releases](https://github.com/tuihub/librarian/releases) page.

### Config file

Config file is required. The config template is provided in [configs](configs)

### Command line arguments

- `config` path, eg: --conf config.yaml
- `data` path, eg: --data /opt/librarian/data

### Environment variables

- `LOG_LEVEL`, accept `debug`, `info`, `warn`, `error`
- `DEMO_MODE`, accept `true`, `false`, server will reject any changes on admin user when demo mode is enabled
- `CREATE_ADMIN_USER`, accept a string
- `CREATE_ADMIN_PASS`, accept a string, server will create the given admin user on startup

## Development

1. Install [Go](https://golang.org/)
2. (Optional) Install [Docker](https://docs.docker.com/) and [Docker Compose](https://docs.docker.com/compose/)
3. (Optional) Deploy dependencies with [tests/docker-compose.yml](tests/docker-compose.yml)
4. Create the config file at [configs/config.yaml](configs/config.yaml) with templates.
5. Run `make run` to start the server
