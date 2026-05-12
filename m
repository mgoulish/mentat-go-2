#! /bin/bash

go mod tidy
go build ./...
go build -o mentat-go-2 ./cmd/mentat-go-2


