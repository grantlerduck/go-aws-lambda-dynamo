#!/usr/bin/env bash

set -eux

echo "Linting code..."
golangci-lint run ./...

echo "Running tests with reports..."
gotestsum --junitfile unit-tests.xml -- -coverprofile=cover.out ./... && \
  go tool cover -html=cover.out -o coverage.html

echo "Checking for vulnerabilities..."
govulncheck ./...

echo "Done"