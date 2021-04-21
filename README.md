## securebanking-openbanking-initialiser
A POC that configures AM and IDM of your CDK deployment, used primarily for testing a sandbox/local dev environment for a securebanking deployment.

## Requirements

- [go 1.15](https://golang.org/doc/install)
- configure [gopath](https://golang.org/doc/gopath_code.html#GOPATH)

## Variables

| Environment variable | description | default |
|----------------------|-------------|---------|
| VERBOSE              | turn on verbose logging | true |
| STRICT               | turn on strict mode, will exit early is invalid statuses are detected | false |
| OPEN_AM_PASSWORD     | The plain text AM password | password |
| REQUEST_BODY_PATH    | path to the directory containing the json requests | config/ |