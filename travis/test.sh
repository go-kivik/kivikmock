#!/bin/bash
set -euC

function join_list {
    local IFS=","
    echo "$*"
}

case "$1" in
    "standard")
        go test -race ./...
    ;;
    "gopherjs")
        gopherjs test ./...
    ;;
    "linter")
        golangci-lint run ./...
    ;;
    "coverage")
        echo "" > coverage.txt

        TEST_PKGS=$(go list ./...)

        for d in $TEST_PKGS; do
            go test -i $d
            go test -coverprofile=profile.out -covermode=set "$d"
            if [ -f profile.out ]; then
                cat profile.out >> coverage.txt
                rm profile.out
            fi
        done

        bash <(curl -s https://codecov.io/bash)
esac
