# Fetch Chainlink Feeds

## Requirements

1. go version v1.12+
2. npx

## Setup

1. Download the dependencies using `go mod download`

## Start

1. Run `make run`

## Notes

1. If you would like to unmarshal the data, you can use the structs provided in [types.go](./types.go)

2. The addresses will be written to a file, `addresses.json`