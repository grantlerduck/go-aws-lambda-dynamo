# go-aws-lambda-dynamo

AWS Lambda function using Go and DynamoDB.
The lambda is implemented in hexagonal architecture (a slight overkill for such a small functions).

Rough functinality & scenario:
* The lambda receives a travel booking event as a message with bas64 encoded payload
* The base64 encoded payload is decoded and marshaled to the protobuf event definition
* The protobuf message is mapped to a domain event
* The domain event is passed to the dynamodb repository and it is inserted into the table
    * In addtion, the repostiroy implements functions to get events by hash and sort key
    * and a global secondary index query
* The repsoitory returns the domain event
* The lambda handler returns the bookingId of the persisted event

This scenario is not particular useful nor fully-real world.
However, it is a nice practice to get the concepts.

Background: I am mainly a Kotlin/Spring Developer and AWS Cloud Architect with a DataSceince backround.
Hence, you will find javaish or pythonish patterns in the code. It has hard to get rid of habbits :D

# AWS Infrastructure

Coming Soon


## Setup
* Install go modules required for local development
* Install gotestsum `go install gotest.tools/gotestsum@latest` (foe test reports in CI tools)
* Install or update ginkgo `go install github.com/onsi/ginkgo/v2/ginkgo`
* Install or update vuln checker `go install golang.org/x/vuln/cmd/govulncheck@latest`
* Install proto compiler `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

## Generate code from protobuf

The booking event lambda receives a json message with a base64 encoded bytestring, the payload, which is parsed to the bookingpb.Event.
The code is generated with protoc and the go compiler plugin as follows:
```bash 
 protoc --go_out=paths=source_relative:. proto/*proto
```

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

### Use ginkgo to bootstrap test suites

make sure you have ginkgo installed, if not install as follows: 
```bash
go install github.com/onsi/ginkgo/v2/ginkgo
```

to bootstrap a new test suite in a module run 
```bash
cd path/to/dir
ginkgo bootstrap
```

Checkout ginkgo [documentation](https://onsi.github.io/ginkgo/) for more details.

## Vulnerability Checks

install vulnerability checker 
```bash 
go install golang.org/x/vuln/cmd/govulncheck@latest
```

run vulnerability check

````bash 
govulncheck ./...
````

## Useful commands

* `cdk deploy`      deploy this stack to your default AWS account/region
* `cdk diff`        compare deployed stack with current state
* `cdk synth`       emits the synthesized CloudFormation template
* `go mod tidy`     remove unused go modules
* `go mod download` install go modules
