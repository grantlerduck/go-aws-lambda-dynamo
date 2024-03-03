# go-aws-lambda-dynamo

AWS Lambda functions using Go, DynamoDB and Kinesis

## Setup
* Install go modules
* Install gotestsum `go install gotest.tools/gotestsum@latest` (fo test reports in CI tools, )

## Testing

Simply run the test script to execute linting and test execution with reports
```bash
./test.sh
```

run unit tests with junit test report and coverage.html report

```bash
gotestsum --junitfile unit-tests.xml -- -coverprofile=cover.out ./... go tool cover -html=cover.out -o coverage.html
```

run linter
````bash
golangci-lint run ./...
````

## Useful commands

* `cdk deploy`      deploy this stack to your default AWS account/region
* `cdk diff`        compare deployed stack with current state
* `cdk synth`       emits the synthesized CloudFormation template
* `go mod tidy`     remove unused go modules
* `go mod download` install go modules
