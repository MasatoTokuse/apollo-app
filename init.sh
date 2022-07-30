#!/bin/zsh
set -e

go mod init $1
printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go
go mod tidy
gqlgen init
gqlgen