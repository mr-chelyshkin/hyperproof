# Hyperproof Take-Home Assessment

## Manual retrieve key:
You can retrieve (or create) any key manually using GitHub manual action: `Manual Retrieve key`.  
Incoming params:
 - Key name for retrieve
 - Targets for key restriction 
 - IPs for key restriction

![manual-retrieve](https://github.com/mr-chelyshkin/hyperproof/blob/main/.img/manual-retrieve.png)  
Execution:  
![manual-retrieve-execution](https://github.com/mr-chelyshkin/hyperproof/blob/main/.img/manual-retrieve-execution.png)

## Schedule retrieve key:
Every 6 hours workflow: `retrieve GoogleMapKey` retrieves a predefined key "GoogleMapKey".  
[retrieve-pipeline](https://github.com/mr-chelyshkin/hyperproof/blob/main/.github/.workflows/schedule-retrieve-GoogleMapKey.yml)
Workflow:
![schedule-flow](https://github.com/mr-chelyshkin/hyperproof/blob/main/.img/schedule-flow.png)
## Client
The project sources is a GoLang client which work with Cloud Platforms (GCP / Azure).  
It's a cli-app based on [urfave](https://github.com/urfave/cli) and use for executing platform actions.

### Usage
```bash
$ hyperproof -h
# Show help with all commands short description
$ hyperproof {{ cmd }} -h
# Show help for current command
```

Command list:
 - retrieve: re-new ApiKey in GCP and put it in Azure Vault.
 - ...

#### retrieve:
```bash
$ hyperproof retrieve -h
```
![retrieve-usage](https://github.com/mr-chelyshkin/hyperproof/blob/main/.img/retrieve-usage.png)
process:
![retrieve-process](https://github.com/mr-chelyshkin/hyperproof/blob/main/.img/retrieve-process.png)
result:
![retrieve-process](https://github.com/mr-chelyshkin/hyperproof/blob/main/.img/retrieve-gcp.png)
![azure-vault-secrets](https://github.com/mr-chelyshkin/hyperproof/blob/main/.img/azure-vault-secrets.png)
![azure-vault-secrets-versions](https://github.com/mr-chelyshkin/hyperproof/blob/main/.img/azure-vault-secrets-versions.png)

### Add new commands
Commands provided in `./commands` directory.
Each cmd is a separate entity located in its directory and consists of the following objects:
 - action.go  - the main process of the team's operation 
 - flags.go   - flags and parameters of the team 
 - usage.go   - template for describing the team's operation 
 - command.go - the main object

To create a new `cmd`, you need to define all the objects and also add the team to commands.go.
```go
return []*cli.Command{
	retrieve.Command(),
        {{ cmd_name }}.Command(),
}
```

## Infrastructure
Use [Taskfile](https://taskfile.dev/) for automate infrastructure actions.
```bash
# current actions:
$ task -l
# task: Available tasks for this project:
# * default:            Default task.
# * go/build/dev:       Build development binary. (include "-race" option)
# * go/build/prd:       Build prd binaries files.
# * go/lint/run:        Run golangci-lint.
```
Inside Taskfile we have predefined variables which we use:
- golint_version
- binary_name

and predefined builds matrix for manage build options:
```yaml
- task: go/build/sample
  vars: { GOOS: "linux", GOARCH: "amd64", BUILD_TYPE: "prd" }
- task: go/build/sample
  vars: { GOOS: "linux", GOARCH: "arm64", BUILD_TYPE: "prd" }
- task: go/build/sample
  vars: { GOOS: "darwin", GOARCH: "amd64", BUILD_TYPE: "prd" }
- task: go/build/sample
  vars: { GOOS: "darwin", GOARCH: "arm64", BUILD_TYPE: "prd" }
```
Build path: `./build`

### Note
- Inside infrastructure workflow we also execute Taskfile.
- By default, we build binaries with GO version from `go.mod` file (_no additional updates need if we decide update version_)

## StaticChecks and Unittests
Any feature addition should be done in separate branches, where the branch name should reflect the general meaning of the feature. 
For each commit to a feature branch, a workflow with general checks is triggered.
![checks-workflow](https://github.com/mr-chelyshkin/hyperproof/blob/main/.img/checks.png)

## Deploy
Deployment is done in a separate workflow, which is triggered after a new tag is published. The deployment process creates a new release, 
builds the client for different platforms, and publishes the artifacts in the release.
![deploy-workflow](https://github.com/mr-chelyshkin/hyperproof/blob/main/.img/deploy.png)

After finishing deployment process new release will be available:
![release](https://github.com/mr-chelyshkin/hyperproof/blob/main/.img/release.png)

For creating new tag:
```bash
# use semver pattern
$ git tag {{ v[0-9]+.\[0-9]+.\[0-9]+.\ }} 
$ git push origin {{ v[0-9]+.\[0-9]+.\[0-9]+.\ }}
```