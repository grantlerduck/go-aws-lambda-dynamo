#!/usr/bin/env bash

set -eux

echo "Linting code..."
golangci-lint run ./...

echo "Running tests with reports..."
gotestsum --junitfile unit-tests.xml -- -coverprofile=cover.out -covermode count ./... && go tool cover -html=cover.out -o coverage.html && gocover-cobertura < cover.out > coverage.xml

echo "Checking for vulnerabilities..."
govulncheck ./...

echo "Syntheizing app..."
cdk synth >> /dev/null

echo "Generating infrastructure diagram..."
npx cdk-dia

echo "Done"