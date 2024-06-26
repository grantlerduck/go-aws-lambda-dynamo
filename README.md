# go-aws-lambda-dynamo - a sample project

AWS Lambda function using Go and DynamoDB.
The lambda is implemented in hexagonal architecture (a slight overkill for such a small function).

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
Hence, you will find javaish or pythonish patterns in the code. It is hard to get rid of habbits :D

## NOTES

* Professionally I have implement various go based ops automation, mainly for CI/CD or similar but I would like to write bigger projects with Go in the future since I really enjoy the simplicity (vs Spring Boot bloated applications)
* AWS CDK for go does not make much sense to me since TypeScript is required anyways and I am used to AWS CDk Typescript for IaC (and having a look at Terraform CDK)
* Go simplicity and performance is a natural fit for serverless functions and container based applications 
* Go seems to be a good replacement for Python which is also commonly used by Platform Enfgineers
* Go is a very eco friendly language (execution time, cpu load, memory consumtpion) while still beeing simple
* Go's ecosystem is developer friendly and the built in dependency management is a dream 
* I was used to CodePipeline and CodeBuild for a long time and V2 is a huge improvement but GitLab CI or GitHub Actions seem to be much simpler and less IaC heavy which is big plus for many developers/orgs (but I am fine with it)
* CodeBuild and CodePipeline can be good if a platform library or automation os provided to reduce the cogntive load for the teams otherwise the alternatives of GitLab and GitHub are just outright better in my opinion
* For alerting I would prefer DataDog as I am used to it professionally but CloudWatch can be good/made easy with [cdk-watchful](https://github.com/cdklabs/cdk-watchful/tree/main) but I did not add it to this sample lambda project

## AWS Infrastructure

Deployed AWS infrastructure using cdk-dia for auto-generating the diagram.
For CICD this sample project uses AWS CodeBuild and CodePipeline using a CodeStarConnection for github.
The pipeline stack contains two pipelines, one for the main branch with execution mode queued and on for specified branches (dev*, feat/* , chore/* , bug/*). 
The pipelines are CodePipelineV2 in order to utilize the new featuers such as branch pipelines, and the different execution modes.  
The stack in this sample assumes a bas setup where the CodeStar connection is available as a stack export to be imported.

![alt text](diagram.png)

### Install CDK-DIA

Globally install cdk-dia 
```
npm install cdk-dia -g
```

Make sure graphviz is installed
```
brew install graphviz
```

Synthesize CDK app
```
cdk synth
```

Generate CDK-DIA diagram as PNG
```
npx cdk-dia
```



## Setup
* Install go modules required for local development
* Install gotestsum `go install gotest.tools/gotestsum@latest` (for test reports in CI tools)
* Install golangci-lint `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest` 
* Install or update ginkgo `go install github.com/onsi/ginkgo/v2/ginkgo`
* Install or update vuln checker `go install golang.org/x/vuln/cmd/govulncheck@latest`
* Install proto compiler `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

## Generate code from protobuf

The booking event lambda receives a json message with a base64 encoded bytestring, the payload, which is parsed to the bookingpb.Event.
The code is generated with protoc and the go compiler plugin as follows:
```bash 
 protoc --go_out=paths=source_relative:. lambda/proto/*proto
```

## Testing

The tests include unit and integrations in a BDD manner.
For the integration tests testcontainers is used to easily automate the container lifetime during test suites.

To execute all tests and reports simply run the test script to execute linting and test execution with reportsTo run
```bash
./test-local.sh
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

## Colima and Testcontainers

If you have a Mac you might be using colima since docker desktop requires a license.
Make sure to correclty configure colima:

```
export TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE=/var/run/docker.sock
export DOCKER_HOST="unix://${HOME}/.colima/docker.sock"
```


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
* `go get -u ./...` update all dependencies recursive 
